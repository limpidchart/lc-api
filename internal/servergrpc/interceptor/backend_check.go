package interceptor

import (
	"context"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/limpidchart/lc-api/internal/backend"
)

// BackendCheck checks if backend is healthy and returns codes.Unavailable status.Status if it's not.
func BackendCheck(log *zerolog.Logger, b backend.Backend) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !b.IsHealthy() {
			log.Error().Msg("Backend connections are not healthy")

			return nil, status.Errorf(codes.Unavailable, "Service Unavailable")
		}

		return handler(ctx, req)
	}
}
