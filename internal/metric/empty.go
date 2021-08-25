package metric

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// EmptyRecorder represents recorder without registered metrics.
type EmptyRecorder struct {
	requestDuration *prometheus.HistogramVec
}

// NewEmptyRecorder returns a new EmptyRecorder.
func NewEmptyRecorder() *EmptyRecorder {
	return &EmptyRecorder{NewRequestDuration()}
}

// RequestDuration returns unregistered request_duration_seconds metric.
func (er *EmptyRecorder) RequestDuration() *prometheus.HistogramVec {
	return er.requestDuration
}

// HTTPHandler returns default Prometheus HTTP handler.
func (er *EmptyRecorder) HTTPHandler() http.Handler {
	return promhttp.Handler()
}
