package middleware

const (
	// RequestIDLogKey represents a request ID key logger field.
	RequestIDLogKey = "request_id"

	// ChartIDLogKey represents a chart ID key logger field.
	ChartIDLogKey = "chart_id"
)

// Context keys.
type ctxKey int

const (
	ctxRequestID ctxKey = iota
	ctxChartID
	ctxCreateChartRequest
)
