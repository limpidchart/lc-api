package servergrpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/limpidchart/lc-api/internal/backend"
	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/metric"
	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/renderer"
	"github.com/limpidchart/lc-api/internal/servergrpc/interceptor"
)

const name = "gRPC API"

// Server implements gRPC render.ChartAPIServer.
type Server struct {
	render.UnimplementedChartAPIServer
	log                *zerolog.Logger
	grpcServer         *grpc.Server
	listener           *net.TCPListener
	rendererClient     render.ChartRendererClient
	rendererReqTimeout time.Duration
	shutdownTimeout    time.Duration
}

// NewServer configures a new Server.
func NewServer(log *zerolog.Logger, tcpList *net.TCPListener, bCon backend.ConnSupervisor, gRPCCfg config.GRPCConfig, pRec metric.PromRecorder) *Server {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.Recover(log),
			interceptor.BackendCheck(log, bCon),
			interceptor.SetRequestID(),
			interceptor.Observer(log, pRec),
		),
	)
	chartAPIServer := &Server{
		log:                log,
		grpcServer:         grpcServer,
		shutdownTimeout:    time.Second * time.Duration(gRPCCfg.ShutdownTimeoutSeconds),
		listener:           tcpList,
		rendererClient:     bCon.RendererClient(),
		rendererReqTimeout: bCon.RendererRequestTimeout(),
	}

	render.RegisterChartAPIServer(grpcServer, chartAPIServer)

	return chartAPIServer
}

// Serve starts gRPC server to serve requests.
func (s *Server) Serve(ctx context.Context) error {
	serveErr := make(chan error)

	// Launch goroutine to serve gRPC requests.
	go func() {
		defer close(serveErr)

		if err := s.grpcServer.Serve(s.listener); err != nil {
			serveErr <- fmt.Errorf("unable to start lc-api gRPC server: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		s.log.Info().
			Time(zerolog.TimestampFieldName, time.Now().UTC()).
			Msg("Trying to gracefully stop lc-api gRPC server")

		stopped := make(chan struct{})

		go func() {
			s.grpcServer.GracefulStop()
			close(stopped)
		}()

		shutdownTimer := time.NewTimer(s.shutdownTimeout)
		select {
		case <-stopped:
			s.log.Info().
				Time(zerolog.TimestampFieldName, time.Now().UTC()).
				Msg("Gracefully stopped lc-api gRPC server")

			shutdownTimer.Stop()

			return nil
		case <-shutdownTimer.C:
			s.log.Warn().
				Time(zerolog.TimestampFieldName, time.Now().UTC()).
				Msg("Unable to gracefully stop lc-api gRPC server, stopping it immediately")

			s.grpcServer.Stop()

			return nil
		}
	case err := <-serveErr:
		return err
	}
}

// Address returns server address.
func (s *Server) Address() string {
	return s.listener.Addr().String()
}

// Name returns server name.
func (s *Server) Name() string {
	return name
}

// CreateChart implements render.ChartAPIServer.CreateChart.
//
// nolint: wrapcheck
func (s *Server) CreateChart(ctx context.Context, req *render.CreateChartRequest) (*render.ChartReply, error) {
	reqID := interceptor.GetRequestID(ctx)

	res, err := renderer.CreateChart(ctx, renderer.CreateChartOpts{
		RequestID:      reqID,
		Request:        req,
		RendererClient: s.rendererClient,
		Timeout:        s.rendererReqTimeout,
	})

	switch {
	case err == nil:
		return res, nil
	case errors.Is(err, renderer.ErrGenerateChartIDFailed):
		return nil, interceptor.InternalError()
	case errors.Is(err, renderer.ErrCreateChartRequestCancelled):
		return nil, status.Error(codes.Canceled, err.Error())
	default:
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
}

// GetChart implements render.ChartAPIServer.GetChart.
// It returns codes.NotFound until storage is implemented.
func (s *Server) GetChart(_ context.Context, req *render.GetChartRequest) (*render.ChartReply, error) {
	return nil, status.Errorf(codes.NotFound, "chart %s is not found", req.ChartId)
}
