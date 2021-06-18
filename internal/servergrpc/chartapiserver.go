package servergrpc

import (
	"context"
	"net"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/limpidchart/lc-api/internal/convert"
	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

// Server implements render.ChartAPIServer.
type Server struct {
	render.UnimplementedChartAPIServer
	grpcServer      *grpc.Server
	listener        *net.TCPListener
	rendererClient  render.ChartRendererClient
	rendererTimeout time.Duration
}

// Opts represents options for new Server.
type Opts struct {
	GRPCServer      *grpc.Server
	Listener        *net.TCPListener
	RendererClient  render.ChartRendererClient
	RendererTimeout time.Duration
}

// NewServer returns a new Server.
func NewServer(opts Opts) *Server {
	//nolint: exhaustivestruct
	chartAPIServer := &Server{
		grpcServer:      opts.GRPCServer,
		listener:        opts.Listener,
		rendererClient:  opts.RendererClient,
		rendererTimeout: opts.RendererTimeout,
	}
	render.RegisterChartAPIServer(opts.GRPCServer, chartAPIServer)

	return chartAPIServer
}

// Address returns listener address.
func (s *Server) Address() string {
	return s.listener.Addr().String()
}

// Serve GRPC server.
func (s *Server) Serve() error {
	//nolint: wrapcheck
	return s.grpcServer.Serve(s.listener)
}

// GracefulStop stops GRPC server gracefully.
func (s *Server) GracefulStop() {
	s.grpcServer.GracefulStop()
}

// CreateChart implements render.ChartAPIServer.CreateChart.
func (s *Server) CreateChart(ctx context.Context, req *render.CreateChartRequest) (*render.ChartReply, error) {
	reqID := uuid.New().String()
	chartID := uuid.New().String()
	now := time.Now().UTC()

	renderChartReq, err := convert.CreateChartRequestToRenderChartRequest(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	renderChartReq.RequestId = reqID

	rendererCtx, rendererCancel := context.WithTimeout(ctx, s.rendererTimeout)
	defer rendererCancel()

	renderReply, renderErr := s.rendererClient.RenderChart(rendererCtx, renderChartReq)
	if renderErr != nil {
		return nil, status.Errorf(codes.InvalidArgument, renderErr.Error())
	}

	return convert.RenderChartReplyToAPIChartReply(reqID, chartID, now, renderReply), nil
}

// GetChart implements render.ChartAPIServer.GetChart.
// It returns codes.NotFound until storage is implemented.
func (s *Server) GetChart(_ context.Context, req *render.GetChartRequest) (*render.ChartReply, error) {
	return nil, status.Errorf(codes.NotFound, "chart %s is not found", req.ChartId)
}
