package servergrpc

import (
	"context"
	"runtime"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

const stackSize = 8192

func recoverInterceptor(log *zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		panicked := true

		defer func() {
			if rec := recover(); rec != nil || panicked {
				stack := make([]byte, stackSize)
				stack = stack[:runtime.Stack(stack, false)]

				log.Error().Bytes(zerolog.ErrorStackFieldName, stack).Interface(zerolog.ErrorFieldName, rec).Msg("gRPC server panicked")

				err = internalError()
			}
		}()

		resp, err := handler(ctx, req)
		panicked = false

		return resp, err
	}
}
