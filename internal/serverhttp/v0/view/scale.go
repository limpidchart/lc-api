package view

// ChartScale represents options to configure chart scale.
// nolint: govet
type ChartScale struct {
	// Kind represents scale kind.
	// Can be one of:
	//  - linear
	//  - band
	//
	// required: true
	Kind string `json:"kind"`

	// RangeStart represents start of the scale range.
	RangeStart *int `json:"range_start,omitempty"`

	// RangeEnd represents end of the scale range.
	RangeEnd *int `json:"range_end,omitempty"`

	// DomainNumeric represents configured numeric domain.
	DomainNumeric *DomainNumeric `json:"domain_numeric,omitempty"`

	// DomainCategories represents configured categories domain.
	DomainCategories *DomainCategories `json:"domain_categories,omitempty"`

	// NoBoundariesOffset disables an offset from the start and end of an axis.
	// This is usually need for an area or line views.
	NoBoundariesOffset bool `json:"no_boundaries_offset,omitempty"`

	// InnerPadding represents inner padding for categories.
	InnerPadding *float32 `json:"inner_padding,omitempty"`

	// OuterPadding represents outer padding for categories.
	OuterPadding *float32 `json:"outer_padding,omitempty"`
}

// DomainNumeric represents numeric scale domain.
type DomainNumeric struct {
	Start float32 `json:"start"`
	End   float32 `json:"end"`
}

// DomainCategories represents string scale domain categories.
type DomainCategories struct {
	Categories []string `json:"categories,omitempty"`
}
