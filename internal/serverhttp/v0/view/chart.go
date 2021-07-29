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
	// Top represents chart top margin.
	Top *int `json:"top,omitempty"`

	// Bottom represents chart bottom margin.
	Bottom *int `json:"bottom,omitempty"`

	// Left represents chart left margin.
	Left *int `json:"left,omitempty"`

	// Right represents chart right margin.
	Right *int `json:"right,omitempty"`
}

// ChartAxes represents options to configure chart axes.
type ChartAxes struct {
	// Top represents configured scale for top axis.
	Top *ChartScale `json:"top,omitempty"`

	// TopLabel represents label for top axis.
	TopLabel string `json:"top_label,omitempty"`

	// Bottom represents configured scale for botom axis.
	Bottom *ChartScale `json:"bottom,omitempty"`

	// BottomLabel represents label for bottom axis.
	BottomLabel string `json:"bottom_label,omitempty"`

	// Left represents configured scale for left axis.
	Left *ChartScale `json:"left,omitempty"`

	// LeftLabel represents label for left axis.
	LeftLabel string `json:"left_label,omitempty"`

	// Right represents configured scale for right axis.
	Right *ChartScale `json:"right,omitempty"`

	// RightLabel represents label for right axis.
	RightLabel string `json:"right_label,omitempty"`
}
