package testutils

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

// TestingChartRendererServer implements render.ChartRendererServer.
//nolint: govet
type TestingChartRendererServer struct {
	render.UnimplementedChartRendererServer
	failMsg    string
	chartData  []byte
	grpcServer *grpc.Server
	listener   *net.TCPListener
}

type RendererServerOpts struct {
	FailMsg   string
	ChartData []byte
}

// NewTestingChartRendererServer returns a new TestingChartRendererServer.
func NewTestingChartRendererServer(opts RendererServerOpts) (*TestingChartRendererServer, error) {
	listener, err := LocalListener()
	if err != nil {
		return nil, fmt.Errorf("unable to prepare local listener: %w", err)
	}

	grpcServer := grpc.NewServer()

	//nolint: exhaustivestruct
	chartRendererServer := &TestingChartRendererServer{
		chartData:  opts.ChartData,
		failMsg:    opts.FailMsg,
		grpcServer: grpcServer,
		listener:   listener,
	}
	render.RegisterChartRendererServer(grpcServer, chartRendererServer)

	return chartRendererServer, nil
}

// Address returns listener address.
func (s *TestingChartRendererServer) Address() string {
	return s.listener.Addr().String()
}

// Serve GRPC server.
func (s *TestingChartRendererServer) Serve() error {
	//nolint: wrapcheck
	return s.grpcServer.Serve(s.listener)
}

// GracefulStop stops GRPC server gracefully.
func (s *TestingChartRendererServer) GracefulStop() {
	s.grpcServer.GracefulStop()
}

// RenderChart implements render.ChartRendererServer.RenderChart.
func (s *TestingChartRendererServer) RenderChart(_ context.Context, req *render.RenderChartRequest) (*render.RenderChartReply, error) {
	if s.failMsg != "" {
		return nil, status.Errorf(codes.InvalidArgument, s.failMsg)
	}

	return &render.RenderChartReply{
		RequestId: req.RequestId,
		ChartData: s.chartData,
	}, nil
}
