package metric

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	requestDurMetricHelp = "The latency of requests (seconds)."
)

// Recorder represents application metrics recorder.
type Recorder interface {
	RequestDuration() *prometheus.HistogramVec
	HTTPHandler() http.Handler
}

type appRecorder struct {
	requestDuration *prometheus.HistogramVec
	registerer      prometheus.Registerer
	httpHandler     http.Handler
}

// NewRecorder registers all metrics and returns a new metric recorder.
// nolint: exhaustivestruct
func NewRecorder() (Recorder, error) {
	registry := prometheus.NewRegistry()

	// Register basic collectors.
	if err := registry.Register(collectors.NewGoCollector()); err != nil {
		return nil, fmt.Errorf("unable to register basic Go collector: %w", err)
	}

	if err := registry.Register(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{})); err != nil {
		return nil, fmt.Errorf("unable to register basic process collector: %w", err)
	}

	// Register custom collectors.
	requestDuration := NewRequestDuration()

	if err := registry.Register(requestDuration); err != nil {
		return nil, fmt.Errorf("unable to register %s metric: %w", requestDurMetricName, err)
	}

	// Configure metrics HTTP handler.
	httpHandler := promhttp.InstrumentMetricHandler(
		registry, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}),
	)

	return &appRecorder{requestDuration, registry, httpHandler}, nil
}

// RequestDuration returns registered request_duration_seconds metric.
func (r *appRecorder) RequestDuration() *prometheus.HistogramVec {
	return r.requestDuration
}

// HTTPHandler returns configured HTTP handler.
func (r *appRecorder) HTTPHandler() http.Handler {
	return r.httpHandler
}

// NewRequestDuration configures and returns a new request_duration_seconds histogram.
func NewRequestDuration() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		// nolint:exhaustivestruct
		prometheus.HistogramOpts{
			Name:    requestDurMetricName,
			Help:    requestDurMetricHelp,
			Buckets: prometheus.DefBuckets,
		},
		[]string{protocolLabel, methodLabel, pathLabel, statusCodeLabel},
	)
}
