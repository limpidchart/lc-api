package view

// BarsValues represents options for bar values.
// nolint: govet
type BarsValues struct {
	// Values contains bars values.
	//
	// required: true
	Values []float32 `json:"values,omitempty"`

	Colors *ChartViewBarsColors `json:"colors,omitempty"`
}

// ChartViewBarsColors represents options to configure bars values colors.
type ChartViewBarsColors struct {
	// Fill represents bars fill color.
	Fill *ChartElementColor `json:"fill,omitempty"`

	// Stroke represents bars stroke color.
	Stroke *ChartElementColor `json:"stroke,omitempty"`
}

// PointsValues represents options for point values.
type PointsValues struct {
	// Values contains points values.
	//
	// required: true
	Values [][]float32 `json:"values,omitempty"`
}

// ScalarValues represents options for scalar values.
type ScalarValues struct {
	// Values contains scalar values.
	//
	// required: true
	Values []float32 `json:"values,omitempty"`
}
