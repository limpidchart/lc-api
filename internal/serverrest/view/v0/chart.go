package view

// ChartSizes represents options to configure chart sizes.
type ChartSizes struct {
	// Width represents chart width.
	Width *int `json:"width,omitempty"`

	// Height represents chart height.
	Height *int `json:"height,omitempty"`
}

// ChartMargins represents options to configure chart margins.
type ChartMargins struct {
	// MarginTop represents chart top margin.
	MarginTop *int `json:"margin_top,omitempty"`

	// MarginBottom represents chart bottom margin.
	MarginBottom *int `json:"margin_bottom,omitempty"`

	// MarginLeft represents chart left margin.
	MarginLeft *int `json:"margin_left,omitempty"`

	// MarginRight represents chart right margin.
	MarginRight *int `json:"margin_right,omitempty"`
}

// ChartAxes represents options to configure chart axes.
type ChartAxes struct {
	// AxisTop represents configured scale for top axis.
	AxisTop *ChartScale `json:"axis_top,omitempty"`

	// AxisTopLabel represents label for top axis.
	AxisTopLabel string `json:"axis_top_label,omitempty"`

	// AxisBottom represents configured scale for botom axis.
	AxisBottom *ChartScale `json:"axis_bottom,omitempty"`

	// AxisBottomLabel represents label for bottom axis.
	AxisBottomLabel string `json:"axis_bottom_label,omitempty"`

	// AxisLeft represents configured scale for left axis.
	AxisLeft *ChartScale `json:"axis_left,omitempty"`

	// AxisLeftLabel represents label for left axis.
	AxisLeftLabel string `json:"axis_left_label,omitempty"`

	// AxisRight represents configured scale for right axis.
	AxisRight *ChartScale `json:"axis_right,omitempty"`

	// AxisRightLabel represents label for right axis.
	AxisRightLabel string `json:"axis_right_label,omitempty"`
}
