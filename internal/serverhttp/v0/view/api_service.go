package view

import "time"

// ChartStatus represents one of available chart statuses.
type ChartStatus string

const (
	// ChartStatusCreated represents a created chart.
	ChartStatusCreated ChartStatus = "CREATED"

	// ChartStatusError represents some error.
	ChartStatusError ChartStatus = "ERROR"
)

func (c ChartStatus) String() string {
	return string(c)
}

// CreateChartRequest represents a request to create chart.
// swagger:parameters createChart
type CreateChartRequest struct {
	// Chart create request body.
	//
	// in: body
	// required: true
	Chart struct {
		// Title represents chart title.
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
	} `json:"chart"`
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
	//
	// swagger:strfmt uuid4
	RequestID string `json:"request_id"`

	// ID of the chart.
	//
	// swagger:strfmt uuid4
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

	// ChartData contains base64 chart representation.
	ChartData string `json:"chart_data"`
}
