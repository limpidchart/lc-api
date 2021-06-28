package servergrpc

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type ctxKey int

const ctxRequestID ctxKey = iota

func requestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		reqID := uuid.New().String()
		newCtx := context.WithValue(ctx, ctxRequestID, reqID)

		return handler(newCtx, req)
	}
}

func getRequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(ctxRequestID).(string); ok {
		return reqID
	}

	return ""
}
