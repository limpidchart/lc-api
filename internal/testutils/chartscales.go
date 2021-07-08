package testutils

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverrest/view/v0"
)

func JSONBandChartScale() *view.ChartScale {
	//nolint: gomnd
	return &view.ChartScale{
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
	}
}

func JSONBandChartScaleTwoDomains() *view.ChartScale {
	res := JSONBandChartScale()

	//nolint: gomnd
	res.DomainNumeric = &view.DomainNumeric{
		Start: 1,
		End:   2,
	}

	return res
}

func BandChartScale() *render.ChartScale {
	//nolint: gomnd
	return &render.ChartScale{
		Kind:       render.ChartScale_BAND,
		RangeStart: &wrapperspb.Int32Value{Value: 0},
		RangeEnd:   &wrapperspb.Int32Value{Value: 100},
		Domain: &render.ChartScale_DomainCategories{
			DomainCategories: &render.DomainCategories{Categories: []string{"A", "B", "C"}},
		},
		NoBoundariesOffset: false,
		InnerPadding:       &wrapperspb.FloatValue{Value: 0.2},
		OuterPadding:       &wrapperspb.FloatValue{Value: 0.2},
	}
}

func JSONBandChartScaleWithNoBoundariesOffset() *view.ChartScale {
	res := JSONBandChartScale()
	res.NoBoundariesOffset = true

	return res
}

func BandChartScaleWithNoBoundariesOffset() *render.ChartScale {
	res := BandChartScale()
	res.NoBoundariesOffset = true

	return res
}

func BandChartScaleWithRanges(rangeStart, rangeEnd int32) *render.ChartScale {
	res := BandChartScale()
	res.RangeStart = &wrapperspb.Int32Value{Value: rangeStart}
	res.RangeEnd = &wrapperspb.Int32Value{Value: rangeEnd}

	return res
}

func BandChartScaleWithoutRanges() *render.ChartScale {
	res := BandChartScale()
	res.RangeStart = nil
	res.RangeEnd = nil

	return res
}

func BandChartScaleWithoutKind() *render.ChartScale {
	res := BandChartScale()
	res.Kind = render.ChartScale_UNSPECIFIED_SCALE

	return res
}

func BandChartScaleWithoutCategoriesDomain() *render.ChartScale {
	res := BandChartScale()

	//nolint: gomnd
	res.Domain = &render.ChartScale_DomainNumeric{
		DomainNumeric: &render.DomainNumeric{
			Start: 20,
			End:   40,
		},
	}

	return res
}

func BandChartScaleWithoutDomain() *render.ChartScale {
	res := BandChartScale()
	res.Domain = nil

	return res
}

func JSONLinearChartScale() *view.ChartScale {
	//nolint: gomnd
	return &view.ChartScale{
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
	}
}

func LinearChartScale() *render.ChartScale {
	//nolint: gomnd
	return &render.ChartScale{
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
	}
}

func LinearChartScaleWithoutRanges() *render.ChartScale {
	res := LinearChartScale()
	res.RangeStart = nil
	res.RangeEnd = nil

	return res
}

func LinearChartScaleWithRangesAndPaddings(rangeStart, rangeEnd int32) *render.ChartScale {
	res := LinearChartScale()
	res.RangeStart = &wrapperspb.Int32Value{Value: rangeStart}
	res.RangeEnd = &wrapperspb.Int32Value{Value: rangeEnd}

	//nolint: gomnd
	res.InnerPadding = &wrapperspb.FloatValue{Value: 0.1}

	//nolint: gomnd
	res.OuterPadding = &wrapperspb.FloatValue{Value: 0.1}

	return res
}

func LinearChartScaleWithDefaults() *render.ChartScale {
	res := LinearChartScale()

	//nolint: gomnd
	res.InnerPadding = &wrapperspb.FloatValue{Value: 0.1}

	//nolint: gomnd
	res.OuterPadding = &wrapperspb.FloatValue{Value: 0.1}

	return res
}

func LinearChartScaleWithDefaultsAndInvertedRanges() *render.ChartScale {
	res := LinearChartScaleWithDefaults()
	res.RangeStart, res.RangeEnd = res.RangeEnd, res.RangeStart

	return res
}

func LinearChartScaleWithoutNumericDomain() *render.ChartScale {
	res := LinearChartScale()
	res.Domain = &render.ChartScale_DomainCategories{
		DomainCategories: &render.DomainCategories{Categories: []string{"a1", "a2"}},
	}

	return res
}
