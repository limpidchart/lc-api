package interceptor

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type ctxKey int

const ctxRequestID ctxKey = iota

// ErrGenerateRequestIDFailed contains error message about failed request ID generation.
var ErrGenerateRequestIDFailed = errors.New("unable to generate a random UUID for chart ID")

// SetRequestID generates and sets request ID into context.
func SetRequestID() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		reqID, err := uuid.NewRandom()
		if err != nil {
			return nil, ErrGenerateRequestIDFailed
		}

		newCtx := context.WithValue(ctx, ctxRequestID, reqID.String())

		return handler(newCtx, req)
	}
}

// GetRequestID returns request ID from context.
func GetRequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(ctxRequestID).(string); ok {
		return reqID
	}

	return ""
}
