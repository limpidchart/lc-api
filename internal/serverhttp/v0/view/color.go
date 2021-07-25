package view

// ChartElementColor represents options to configure color for chart elements.
//nolint: govet
type ChartElementColor struct {
	// Hex represents hex color value.
	Hex string `json:"hex,omitempty"`

	// RGB represents RGB color value.
	RGB *RGB `json:"rgb,omitempty"`
}

// RGB contains values for RGB color.
type RGB struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}
