package servergrpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func recoverInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		panicked := true

		defer func() {
			if rec := recover(); rec != nil || panicked {
				err = status.Error(codes.Internal, "Unable to handle request")
			}
		}()

		resp, err := handler(ctx, req)
		panicked = false

		return resp, err
	}
}
