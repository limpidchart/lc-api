package servergrpc

import (
	"context"
	"path"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

const (
	requestIDKey = "request_id"
	ipKey        = "ip"
	codeKey      = "code"
	methodKey    = "method"
	durationKey  = "duration"
	errKey       = "error"
)

const unknownIP = "unknown"

func loggerInterceptor(log *zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now().UTC()
		reqID := getRequestID(ctx)

		resp, err := handler(ctx, req)
		if err != nil {
			logEvent := basicLoggerFields(ctx, log.Error(), startTime, reqID, info.FullMethod, err)
			logEvent = errLoggerFields(logEvent, err)
			logEvent.Msg("")

			return resp, err
		}

		logEvent := basicLoggerFields(ctx, log.Info(), startTime, reqID, info.FullMethod, err)
		logEvent.Msg("")

		return resp, err
	}
}

func basicLoggerFields(ctx context.Context, logEvent *zerolog.Event, startTime time.Time, reqID, method string, err error) *zerolog.Event {
	return logEvent.
		Time(zerolog.TimestampFieldName, startTime).
		Str(requestIDKey, reqID).
		Str(ipKey, peerIP(ctx)).
		Str(codeKey, status.Convert(err).Code().String()).
		Str(methodKey, path.Base(method)).
		Dur(durationKey, time.Since(startTime))
}

func errLoggerFields(logEvent *zerolog.Event, err error) *zerolog.Event {
	return logEvent.Str(errKey, status.Convert(err).Message())
}

func peerIP(ctx context.Context) string {
	ip, ok := peer.FromContext(ctx)
	if ok {
		return ip.Addr.String()
	}

	return unknownIP
}
