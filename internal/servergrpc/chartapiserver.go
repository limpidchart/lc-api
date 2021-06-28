package servergrpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/convert"
	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/tcputils"
)

const rendererServiceCfg = `{"loadBalancingPolicy":"round_robin"}`

// ErrCreateChartRequestCancelled contains error message about cancelled create chart request.
var ErrCreateChartRequestCancelled = errors.New("create chart request is cancelled")

// Server implements render.ChartAPIServer.
type Server struct {
	render.UnimplementedChartAPIServer
	log                *zerolog.Logger
	grpcServer         *grpc.Server
	listener           *net.TCPListener
	rendererConn       *grpc.ClientConn
	rendererClient     render.ChartRendererClient
	rendererReqTimeout time.Duration
	shutdownTimeout    time.Duration
}

// NewServer configures needed connections and returns a new Server.
func NewServer(ctx context.Context, log *zerolog.Logger, apiCfg config.APIConfig, rendererCfg config.RendererConfig) (*Server, error) {
	listener, err := tcputils.Listener(apiCfg.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to start lc-api TCP listener: %w", err)
	}

	rendererConnCtx, rendererConnCancel := context.WithTimeout(ctx, time.Second*time.Duration(rendererCfg.ConnTimeoutSeconds))
	defer rendererConnCancel()

	rendererConn, err := grpc.DialContext(
		rendererConnCtx,
		rendererCfg.Address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(rendererServiceCfg),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection to lc-renderer: %w", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(recoverInterceptor(), requestIDInterceptor(), loggerInterceptor(log)),
	)

	//nolint: exhaustivestruct
	chartAPIServer := &Server{
		log:                log,
		grpcServer:         grpcServer,
		shutdownTimeout:    time.Second * time.Duration(apiCfg.ShutdownTimeoutSeconds),
		listener:           listener,
		rendererConn:       rendererConn,
		rendererClient:     render.NewChartRendererClient(rendererConn),
		rendererReqTimeout: time.Second * time.Duration(rendererCfg.RequestTimeoutSeconds),
	}

	render.RegisterChartAPIServer(grpcServer, chartAPIServer)

	return chartAPIServer, nil
}

// Address returns listener address.
func (s *Server) Address() string {
	return s.listener.Addr().String()
}

// Serve start gRPC server to serve requests.
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
			s.rendererConn.Close()
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

// CreateChart implements render.ChartAPIServer.CreateChart.
func (s *Server) CreateChart(ctx context.Context, req *render.CreateChartRequest) (*render.ChartReply, error) {
	reqID := getRequestID(ctx)
	chartID := uuid.New().String()
	now := time.Now().UTC()

	renderChartReq, err := convert.CreateChartRequestToRenderChartRequest(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	renderChartReq.RequestId = reqID

	rendererCtx, rendererCancel := context.WithTimeout(ctx, s.rendererReqTimeout)
	defer rendererCancel()

	select {
	case <-ctx.Done():
		return nil, ErrCreateChartRequestCancelled
	case renderResult := <-s.renderChart(rendererCtx, renderChartReq):
		if renderResult.err != nil {
			return nil, status.Errorf(codes.InvalidArgument, renderResult.err.Error())
		}

		return convert.RenderChartReplyToAPIChartReply(reqID, chartID, now, renderResult.reply), nil
	}
}

type renderChartResult struct {
	reply *render.RenderChartReply
	err   error
}

func (s *Server) renderChart(ctx context.Context, req *render.RenderChartRequest) <-chan renderChartResult {
	result := make(chan renderChartResult)

	go func() {
		reply, err := s.rendererClient.RenderChart(ctx, req)
		result <- renderChartResult{reply, err}
		close(result)
	}()

	return result
}

// GetChart implements render.ChartAPIServer.GetChart.
// It returns codes.NotFound until storage is implemented.
func (s *Server) GetChart(_ context.Context, req *render.GetChartRequest) (*render.ChartReply, error) {
	return nil, status.Errorf(codes.NotFound, "chart %s is not found", req.ChartId)
}
