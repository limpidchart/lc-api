package interceptor

import (
	"context"
	"path"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/limpidchart/lc-api/internal/metric"
)

const (
	protocolKey  = "protocol"
	requestIDKey = "request_id"
	ipKey        = "ip"
	codeKey      = "code"
	methodKey    = "method"
	durationKey  = "duration"
	errKey       = "error"
)

const unknownIP = "unknown"

// Observer handles observability (metrics and logging) for every request.
func Observer(log *zerolog.Logger, mrec metric.Recorder) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now().UTC()

		resp, err := handler(ctx, req)

		observe(ctx, log, startTime, info, mrec, err)

		return resp, err
	}
}

func observe(ctx context.Context, log *zerolog.Logger, startTime time.Time, info *grpc.UnaryServerInfo, mrec metric.Recorder, err error) {
	reqID := GetRequestID(ctx)
	duration := time.Since(startTime)

	if err != nil {
		logEvent := basicLoggerFields(ctx, log.Error(), startTime, duration, reqID, info.FullMethod, err)
		logEvent = errLoggerFields(logEvent, err)
		logEvent.Msg("")
	} else {
		logEvent := basicLoggerFields(ctx, log.Info(), startTime, duration, reqID, info.FullMethod, err)
		logEvent.Msg("")
	}

	mrec.RequestDuration().WithLabelValues(metric.ProtocolGRPC, path.Base(info.FullMethod), info.FullMethod, status.Convert(err).Code().String()).Observe(duration.Seconds())
}

func basicLoggerFields(ctx context.Context, logEvent *zerolog.Event, startTime time.Time, duration time.Duration, reqID, method string, err error) *zerolog.Event {
	return logEvent.
		Time(zerolog.TimestampFieldName, startTime).
		Str(protocolKey, metric.ProtocolGRPC).
		Str(requestIDKey, reqID).
		Str(ipKey, peerIP(ctx)).
		Str(codeKey, status.Convert(err).Code().String()).
		Str(methodKey, path.Base(method)).
		Dur(durationKey, duration)
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
