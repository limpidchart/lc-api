package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// ProtocolHTTP represents protocol name for HTTP requests.
	ProtocolHTTP = "http"

	// ProtocolGRPC represents protocol name for gRPC requests.
	ProtocolGRPC = "grpc"
)

const (
	protocolLabel   = "protocol"
	methodLabel     = "method"
	pathLabel       = "path"
	statusCodeLabel = "status_code"

	requestDurMetricName = "request_duration_seconds"
	requestDurMetricHelp = "The latency of requests."
)

// RequestDuration represents a metrics with duration of each request.
func RequestDuration() *prometheus.HistogramVec {
	// nolint: exhaustivestruct
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    requestDurMetricName,
			Help:    requestDurMetricHelp,
			Buckets: prometheus.DefBuckets,
		},
		[]string{protocolLabel, methodLabel, pathLabel, statusCodeLabel},
	)
}
