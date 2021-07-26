package view

const ParamChartID = "chart_id"

// ChartID represents chart ID from URL.
//
// swagger:parameters getChart
type ChartID struct {
	// Chart identifier.
	//
	// in: path
	// type: string
	ChartID string `json:"chart_id"`
}
