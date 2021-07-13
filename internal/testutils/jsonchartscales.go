package testutils

import "github.com/limpidchart/lc-api/internal/serverrest/view/v0"

type JSONChartScale struct {
	*view.ChartScale
}

func NewJSONBandChartScale() *JSONChartScale {
	// nolint: gomnd
	return &JSONChartScale{
		&view.ChartScale{
			Kind:          "band",
			RangeStart:    intToPtr(0),
			RangeEnd:      intToPtr(100),
			DomainNumeric: nil,
			DomainCategories: &view.DomainCategories{
				Categories: []string{"A", "B", "C"},
			},
			NoBoundariesOffset: false,
			InnerPadding:       float32ToPtr(0.2),
			OuterPadding:       float32ToPtr(0.2),
		},
	}
}

func NewJSONLinearChartScale() *JSONChartScale {
	// nolint: gomnd
	return &JSONChartScale{
		&view.ChartScale{
			Kind:       "linear",
			RangeStart: intToPtr(11),
			RangeEnd:   intToPtr(111),
			DomainNumeric: &view.DomainNumeric{
				Start: 0,
				End:   100,
			},
			DomainCategories:   nil,
			NoBoundariesOffset: false,
			InnerPadding:       nil,
			OuterPadding:       nil,
		},
	}
}

func (s *JSONChartScale) Unembed() *view.ChartScale {
	return s.ChartScale
}

func (s *JSONChartScale) SetTwoDomains() *JSONChartScale {
	// nolint: gomnd
	s.DomainNumeric = &view.DomainNumeric{
		Start: 1,
		End:   2,
	}

	return s
}

func (s *JSONChartScale) SetNoBoundariesOffset() *JSONChartScale {
	s.NoBoundariesOffset = true

	return s
}
