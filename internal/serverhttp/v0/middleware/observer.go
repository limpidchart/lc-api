package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"

	"github.com/limpidchart/lc-api/internal/metric"
)

const (
	xRealIPHeader       = "x-real-ip"
	xForwardedForHeader = "x-forwarded-for"
	userAgentHeader     = "user-agent"
	refererHeader       = "referer"
)

const (
	protocolKey     = "protocol"
	ipKey           = "ip"
	userAgentKey    = "user_agent"
	refererKey      = refererHeader
	codeKey         = "code"
	methodKey       = "method"
	pathKey         = "path"
	bytesWrittenKey = "resp_bytes_written"
	durationKey     = "duration"
)

const (
	errCodesStart  = 500
	warnCodesStart = 400
)

// RequestObserver handles observability (metrics and logging) for every request.
func RequestObserver(log *zerolog.Logger, pRec metric.PromRecorder) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := chimiddleware.NewWrapResponseWriter(w, r.ProtoMajor)

			startTime := time.Now().UTC()

			next.ServeHTTP(ww, r)

			defer func() {
				statusCode := ww.Status()
				bytesWritten := ww.BytesWritten()
				duration := time.Since(startTime)

				pRec.RequestDuration().WithLabelValues(metric.ProtocolHTTP, r.Method, r.URL.Path, strconv.Itoa(statusCode)).Observe(duration.Seconds())

				switch {
				case statusCode >= errCodesStart:
					logEvent := loggerFields(log.Error(), r, statusCode, bytesWritten, startTime, duration)
					logEvent.Msg("")
				case statusCode >= warnCodesStart:
					logEvent := loggerFields(log.Warn(), r, statusCode, bytesWritten, startTime, duration)
					logEvent.Msg("")
				default:
					logEvent := loggerFields(log.Info(), r, statusCode, bytesWritten, startTime, duration)
					logEvent.Msg("")
				}
			}()
		})
	}
}

func loggerFields(logEvent *zerolog.Event, r *http.Request, code, bytesWritten int, startTime time.Time, duration time.Duration) *zerolog.Event {
	event := logEvent.
		Time(zerolog.TimestampFieldName, startTime).
		Str(protocolKey, metric.ProtocolHTTP).
		Str(RequestIDLogKey, GetRequestID(r.Context())).
		Str(ipKey, peerIP(r)).
		Str(userAgentKey, r.Header.Get(userAgentHeader)).
		Str(refererKey, r.Header.Get(refererHeader)).
		Int(codeKey, code).
		Str(methodKey, r.Method).
		Str(pathKey, r.URL.Path).
		Int(bytesWrittenKey, bytesWritten).
		Dur(durationKey, duration)

	chartID := GetChartID(r.Context())
	if chartID == "" {
		return event
	}

	return event.Str(ChartIDLogKey, GetChartID(r.Context()))
}

func peerIP(r *http.Request) string {
	realIP := r.Header.Get(xRealIPHeader)
	forwardedFor := r.Header.Get(xForwardedForHeader)

	if realIP == "" && forwardedFor == "" {
		return ipFromRemoteAddr(r.RemoteAddr)
	}

	if forwardedFor == "" {
		return realIP
	}

	return strings.Split(forwardedFor, ",")[0]
}

func ipFromRemoteAddr(s string) string {
	idx := strings.LastIndex(s, ":")
	if idx == -1 {
		return s
	}

	return s[:idx]
}
