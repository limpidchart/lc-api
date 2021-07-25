package view

// BarsValues represents options for bar values.
//nolint: govet
type BarsValues struct {
	// Values contains bars values.
	//
	// required: true
	Values []float32 `json:"values,omitempty"`

	// FillColor represents bars fill color.
	FillColor *ChartElementColor `json:"fill_color,omitempty"`

	// StrokeColor represents bars stroke color.
	StrokeColor *ChartElementColor `json:"stroke_color,omitempty"`
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
