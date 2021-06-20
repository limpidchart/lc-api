package testutils

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/tcputils"
)

// ErrRequestCancelled contains error message about cancelled testing lc-renderer request.
var ErrRequestCancelled = errors.New("request to testing lc-renderer is cancelled")

// TestingChartRendererServer implements render.ChartRendererServer.
type TestingChartRendererServer struct {
	render.UnimplementedChartRendererServer
	failMsg    string
	grpcServer *grpc.Server
	listener   *net.TCPListener
	chartData  []byte
	latency    time.Duration
}

// Opts contains options to configure TestingChartRendererServer.
type Opts struct {
	FailMsg   string
	ChartData []byte
	Latency   time.Duration
}

// NewTestingChartRendererServer returns a new TestingChartRendererServer.
func NewTestingChartRendererServer(opts Opts) (*TestingChartRendererServer, error) {
	listener, err := tcputils.LocalListenerWithRandomPort()
	if err != nil {
		return nil, fmt.Errorf("unable to prepare local listener: %w", err)
	}

	grpcServer := grpc.NewServer()

	//nolint: exhaustivestruct
	chartRendererServer := &TestingChartRendererServer{
		failMsg:    opts.FailMsg,
		grpcServer: grpcServer,
		listener:   listener,
		chartData:  opts.ChartData,
		latency:    opts.Latency,
	}

	render.RegisterChartRendererServer(grpcServer, chartRendererServer)

	return chartRendererServer, nil
}

// Address returns listener address.
func (s *TestingChartRendererServer) Address() string {
	return s.listener.Addr().String()
}

// Serve gRPC server.
func (s *TestingChartRendererServer) Serve(ctx context.Context) error {
	serveErr := make(chan error)

	// Launch goroutine to serve gRPC requests.
	go func() {
		defer close(serveErr)

		if err := s.grpcServer.Serve(s.listener); err != nil {
			serveErr <- fmt.Errorf("unable to start testing lc-renderer gRPC server: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		s.grpcServer.GracefulStop()

		return nil
	case err := <-serveErr:
		return err
	}
}

// RenderChart implements render.ChartRendererServer.RenderChart.
func (s *TestingChartRendererServer) RenderChart(ctx context.Context, req *render.RenderChartRequest) (*render.RenderChartReply, error) {
	// Render chart with the provided latency.
	renderTimer := time.NewTimer(s.latency)

	select {
	case <-ctx.Done():
		return nil, ErrRequestCancelled
	case <-renderTimer.C:
		if s.failMsg != "" {
			return nil, status.Errorf(codes.InvalidArgument, s.failMsg)
		}

		return &render.RenderChartReply{
			RequestId: req.RequestId,
			ChartData: s.chartData,
		}, nil
	}
}
