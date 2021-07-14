package testutils

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

type ChartScale struct {
	*render.ChartScale
}

func NewBandChartScale() *ChartScale {
	// nolint: gomnd
	return &ChartScale{
		&render.ChartScale{
			Kind:       render.ChartScale_BAND,
			RangeStart: &wrapperspb.Int32Value{Value: 0},
			RangeEnd:   &wrapperspb.Int32Value{Value: 100},
			Domain: &render.ChartScale_DomainCategories{
				DomainCategories: &render.DomainCategories{Categories: []string{"A", "B", "C"}},
			},
			NoBoundariesOffset: false,
			InnerPadding:       &wrapperspb.FloatValue{Value: 0.2},
			OuterPadding:       &wrapperspb.FloatValue{Value: 0.2},
		},
	}
}

func NewLinearChartScale() *ChartScale {
	// nolint: gomnd
	return &ChartScale{
		&render.ChartScale{
			Kind:       render.ChartScale_LINEAR,
			RangeStart: &wrapperspb.Int32Value{Value: 11},
			RangeEnd:   &wrapperspb.Int32Value{Value: 111},
			Domain: &render.ChartScale_DomainNumeric{
				DomainNumeric: &render.DomainNumeric{
					Start: 0,
					End:   100,
				},
			},
			NoBoundariesOffset: false,
			InnerPadding:       nil,
			OuterPadding:       nil,
		},
	}
}

func (s *ChartScale) Unembed() *render.ChartScale {
	return s.ChartScale
}

func (s *ChartScale) SetNoBoundariesOffset() *ChartScale {
	s.NoBoundariesOffset = true

	return s
}

func (s *ChartScale) SetRanges(rangeStart, rangeEnd int32) *ChartScale {
	s.RangeStart = &wrapperspb.Int32Value{Value: rangeStart}
	s.RangeEnd = &wrapperspb.Int32Value{Value: rangeEnd}

	return s
}

func (s *ChartScale) UnsetRanges() *ChartScale {
	s.RangeStart = nil
	s.RangeEnd = nil

	return s
}

func (s *ChartScale) UnsetKind() *ChartScale {
	s.Kind = render.ChartScale_UNSPECIFIED_SCALE

	return s
}

func (s *ChartScale) SetNumericDomain() *ChartScale {
	// nolint: gomnd
	s.Domain = &render.ChartScale_DomainNumeric{
		DomainNumeric: &render.DomainNumeric{
			Start: 20,
			End:   40,
		},
	}

	return s
}

func (s *ChartScale) UnsetDomain() *ChartScale {
	s.Domain = nil

	return s
}

func (s *ChartScale) SetPaddings() *ChartScale {
	// nolint: gomnd
	s.InnerPadding = &wrapperspb.FloatValue{Value: 0.1}

	// nolint: gomnd
	s.OuterPadding = &wrapperspb.FloatValue{Value: 0.1}

	return s
}

func (s *ChartScale) InvertRanges() *ChartScale {
	s.RangeStart, s.RangeEnd = s.RangeEnd, s.RangeStart

	return s
}

func (s *ChartScale) SetCategoriesDomain() *ChartScale {
	s.Domain = &render.ChartScale_DomainCategories{
		DomainCategories: &render.DomainCategories{Categories: []string{"a1", "a2"}},
	}

	return s
}
