package servergrpchealthcheck

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/limpidchart/lc-api/internal/backend"
	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/tcputils"
)

// Server implements gRPC grpc_health_v1.HealthServer.
type Server struct {
	grpc_health_v1.UnimplementedHealthServer
	log        *zerolog.Logger
	grpcServer *grpc.Server
	listener   *net.TCPListener
	bCon       backend.ConnSupervisor
}

// NewServer configures a new Server.
func NewServer(log *zerolog.Logger, bCon backend.ConnSupervisor, hcCfg config.GRPCHealthCheckConfig) (*Server, error) {
	listener, err := tcputils.Listener(hcCfg.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to start lc-api gRPC health check listener: %w", err)
	}

	grpcServer := grpc.NewServer()
	hcServer := &Server{
		log:        log,
		grpcServer: grpcServer,
		listener:   listener,
		bCon:       bCon,
	}

	grpc_health_v1.RegisterHealthServer(grpcServer, hcServer)

	return hcServer, nil
}

// Serve starts gRPC health check server to serve requests.
func (s *Server) Serve(ctx context.Context) error {
	serveErr := make(chan error)

	// Launch goroutine to serve gRPC requests.
	go func() {
		defer close(serveErr)

		if err := s.grpcServer.Serve(s.listener); err != nil {
			serveErr <- fmt.Errorf("unable to start lc-api gRPC health check server: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		s.log.Info().
			Time(zerolog.TimestampFieldName, time.Now().UTC()).
			Msg("Stopping lc-api gRPC health check server")

		s.grpcServer.Stop()

		return nil
	case err := <-serveErr:
		return err
	}
}

// Address returns listener address.
func (s *Server) Address() string {
	return s.listener.Addr().String()
}

// Check implements grpc_health_v1.HealthServer.Check.
func (s *Server) Check(_ context.Context, _ *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	status := grpc_health_v1.HealthCheckResponse_SERVING

	if !s.bCon.IsHealthy() {
		status = grpc_health_v1.HealthCheckResponse_NOT_SERVING
	}

	return &grpc_health_v1.HealthCheckResponse{
		Status: status,
	}, nil
}
