package testutils

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

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

func LinearChartScale() *render.ChartScale {
	//nolint: gomnd
	return &render.ChartScale{
		Kind:       render.ChartScale_LINEAR,
		RangeStart: &wrapperspb.Int32Value{Value: 0},
		RangeEnd:   &wrapperspb.Int32Value{Value: 100},
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

func LinearChartScaleWithDefaults() *render.ChartScale {
	res := LinearChartScale()

	//nolint: gomnd
	res.InnerPadding = &wrapperspb.FloatValue{Value: 0.1}

	//nolint: gomnd
	res.OuterPadding = &wrapperspb.FloatValue{Value: 0.1}

	return res
}

func LinearChartScaleWithoutNumericDomain() *render.ChartScale {
	res := LinearChartScale()
	res.Domain = &render.ChartScale_DomainCategories{
		DomainCategories: &render.DomainCategories{Categories: []string{"a1", "a2"}},
	}

	return res
}
