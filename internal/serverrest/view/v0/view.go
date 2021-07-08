package view

// ChartView represents options to configure chart view.
//nolint: govet
type ChartView struct {
	// Kind represents view kind.
	// Can be one of:
	//  - area
	//  - horizontal_bar
	//  - line
	//  - scatter
	//  - vertical_bar
	//
	// required: true
	Kind string `json:"kind"`

	// BarsValues represents bars values.
	// It can be used with horizontal or vertical bar view.
	BarsValues []*BarsValues `json:"bars_values,omitempty"`

	// PointsValues represents points values.
	// It can be used with scatter view only.
	PointsValues *PointsValues `json:"points_values,omitempty"`

	// ScalarValues represents scalar values.
	// It can be used with area or line view.
	ScalarValues *ScalarValues `json:"scalar_values,omitempty"`

	// Colors represents view colors.
	Colors *ChartViewColors `json:"colors,omitempty"`

	// BarLabelVisible represents bar visibility.
	BarLabelVisible *bool `json:"bar_label_visible,omitempty"`

	// BarLabelPosition represents bar label position.
	// Can be one of:
	//  - start_outside
	//  - start_inside
	//  - center
	//  - end_inside
	//  - end_outside
	BarLabelPosition string `json:"bar_label_position,omitempty"`

	// PointVisible represents point visibility.
	PointVisible *bool `json:"point_visible,omitempty"`

	// PointType represents view point type.
	// Can be one of:
	//  - circle
	//  - square
	//  - x
	PointType string `json:"point_type,omitempty"`

	// PointLabelVisible represents point visibility.
	PointLabelVisible *bool `json:"point_label_visible,omitempty"`

	// PointLabelPosition represents point label position.
	// Can be one of:
	//  - top
	//  - top_right
	//  - top_left
	//  - left
	//  - right
	//  - bottom
	//  - bottom_left
	//  - bottom_right
	PointLabelPosition string `json:"point_label_position,omitempty"`
}

// ChartViewColors represents view colors parameters.
type ChartViewColors struct {
	// FillColor represents view fill color.
	FillColor *ChartElementColor `json:"fill_color,omitempty"`

	// StrokeColor represents view stroke color.
	StrokeColor *ChartElementColor `json:"stroke_color,omitempty"`

	// PointFillColor represents view point fill color.
	PointFillColor *ChartElementColor `json:"point_fill_color,omitempty"`

	// PointStrokeColor represents view point stroke color.
	PointStrokeColor *ChartElementColor `json:"point_stroke_color,omitempty"`
}
