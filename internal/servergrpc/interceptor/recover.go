package interceptor

import (
	"context"
	"runtime"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const stackSize = 8192

// Recover represents interceptro to catch panics.
func Recover(log *zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		panicked := true

		defer func() {
			if rec := recover(); rec != nil || panicked {
				stack := make([]byte, stackSize)
				stack = stack[:runtime.Stack(stack, false)]

				log.Error().Bytes(zerolog.ErrorStackFieldName, stack).Interface(zerolog.ErrorFieldName, rec).Msg("gRPC server panicked")

				err = InternalError()
			}
		}()

		resp, err := handler(ctx, req)
		panicked = false

		return resp, err
	}
}

// InternalError returns status.Error with codes.Internal code.
func InternalError() error {
	// nolint: wrapcheck
	return status.Error(codes.Internal, "server is unable to handle the request")
}
