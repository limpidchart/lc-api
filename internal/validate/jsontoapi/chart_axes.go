package jsontoapi

import (
	"errors"
	"strings"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
)

const (
	scaleLinearKind = "linear"
	scaleBandKind   = "band"
)

// ErrOnlyOneOfDomainsShouldBeSpecified contains error message about a case when too many domains are provided.
var ErrOnlyOneOfDomainsShouldBeSpecified = errors.New("only one of domain_numeric or domain_categories should be specified")

// ChartScaleFromJSON parses and validates JSON chart scale representation.
func ChartScaleFromJSON(scale *view.ChartScale) (*render.ChartScale, error) {
	if scale == nil {
		return nil, nil
	}

	if scale.DomainNumeric != nil && scale.DomainCategories != nil {
		return nil, ErrOnlyOneOfDomainsShouldBeSpecified
	}

	result := emptyChartScale()

	if scale.DomainNumeric != nil {
		result.Domain = &render.ChartScale_DomainNumeric{
			DomainNumeric: &render.DomainNumeric{
				Start: scale.DomainNumeric.Start,
				End:   scale.DomainNumeric.End,
			},
		}
	}

	if scale.DomainCategories != nil {
		result.Domain = &render.ChartScale_DomainCategories{
			DomainCategories: &render.DomainCategories{
				Categories: scale.DomainCategories.Categories,
			},
		}
	}

	var (
		rangeStart   *wrapperspb.Int32Value
		rangeEnd     *wrapperspb.Int32Value
		innerPadding *wrapperspb.FloatValue
		outerPadding *wrapperspb.FloatValue
	)

	if scale.RangeStart != nil {
		rangeStart = &wrapperspb.Int32Value{Value: int32(*scale.RangeStart)}
	}

	if scale.RangeEnd != nil {
		rangeEnd = &wrapperspb.Int32Value{Value: int32(*scale.RangeEnd)}
	}

	if scale.InnerPadding != nil {
		innerPadding = &wrapperspb.FloatValue{Value: *scale.InnerPadding}
	}

	if scale.OuterPadding != nil {
		outerPadding = &wrapperspb.FloatValue{Value: *scale.OuterPadding}
	}

	result.Kind = scaleKindFromJSON(scale.Kind)
	result.RangeStart = rangeStart
	result.RangeEnd = rangeEnd
	result.NoBoundariesOffset = scale.NoBoundariesOffset
	result.InnerPadding = innerPadding
	result.OuterPadding = outerPadding

	return result, nil
}

func emptyChartScale() *render.ChartScale {
	return &render.ChartScale{
		Kind:               0,
		RangeStart:         nil,
		RangeEnd:           nil,
		Domain:             nil,
		NoBoundariesOffset: false,
		InnerPadding:       nil,
		OuterPadding:       nil,
	}
}

func scaleKindFromJSON(kind string) render.ChartScale_ChartScaleKind {
	switch strings.ToLower(kind) {
	case scaleLinearKind:
		return render.ChartScale_LINEAR
	case scaleBandKind:
		return render.ChartScale_BAND
	default:
		return render.ChartScale_UNSPECIFIED_SCALE
	}
}
