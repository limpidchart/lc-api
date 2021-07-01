package view

import "time"

// CreateChartRequest represents a request to create chart.
// swagger:parameters createChart
type CreateChartRequest struct {
	// Chart create request body.
	//
	// in: body
	// required: true
	Request struct {
		// Title represents chart title.
		//
		// required: true
		Title string `json:"title"`

		// Sizes represents chart sizes.
		Sizes *ChartSizes `json:"sizes"`

		// Margins represents chart margins.
		Margins *ChartMargins `json:"margins"`

		// Axes represents chart axes.
		//
		// required: true
		Axes *ChartAxes `json:"axes"`

		// Views represents chart views
		//
		// required: true
		Views []*ChartView `json:"views"`
	}
}

// GetChartRequest represents a request to get chart.
// swagger:parameters getChart
type GetChartRequest struct {
	// ChartID represents id of the chart.
	//
	// required: true
	// swagger:strfmt uuid4
	ChartID string `json:"chat_id"`
}

// ChartReply represents a reply from create or get requests.
// swagger:model chartReply
type ChartReply struct {
	// ID of the request.
	RequestID string `json:"request_id"`

	// ID of the chart.
	ChartID string `json:"chart_id"`

	// Chart status.
	// Can be one of:
	//  - CREATED
	//  - ERROR
	ChartStatus string `json:"chart_status"`

	// CreatedAt contains chart creation timestamp.
	CreatedAt *time.Time `json:"created_at"`

	// DeletedAt contains chart deletion timestamp.
	DeletedAt *time.Time `json:"deleted_at"`

	// ChartData represents chart raw bytes representation.
	ChartData []byte `json:"chart_data"`
}
