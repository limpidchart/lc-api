package apitorenderer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/validate/apitorenderer"
)

func testingLinearChartScale() *render.ChartScale {
	return &render.ChartScale{
		Kind:       render.ChartScale_LINEAR,
		RangeStart: &wrapperspb.Int32Value{Value: 10},
		RangeEnd:   &wrapperspb.Int32Value{Value: 100},
		Domain: &render.ChartScale_DomainNumeric{
			DomainNumeric: &render.DomainNumeric{
				Start: 20,
				End:   40,
			},
		},
		NoBoundariesOffset: false,
		InnerPadding:       nil,
		OuterPadding:       nil,
	}
}

func testingBandChartScale() *render.ChartScale {
	return &render.ChartScale{
		Kind:       render.ChartScale_BAND,
		RangeStart: &wrapperspb.Int32Value{Value: 0},
		RangeEnd:   &wrapperspb.Int32Value{Value: 1000},
		Domain: &render.ChartScale_DomainCategories{
			DomainCategories: &render.DomainCategories{Categories: []string{"a1", "a2"}},
		},
		NoBoundariesOffset: true,
		InnerPadding:       &wrapperspb.FloatValue{Value: 0.5},
		OuterPadding:       &wrapperspb.FloatValue{Value: 0.5},
	}
}

func TestValidateChartAxes(t *testing.T) {
	t.Parallel()

	testingLinearChartScaleWithPaddings := testingLinearChartScale()
	testingLinearChartScaleWithPaddings.InnerPadding = &wrapperspb.FloatValue{Value: 0.1}
	testingLinearChartScaleWithPaddings.OuterPadding = &wrapperspb.FloatValue{Value: 0.1}

	testingScaleWithoutKind := testingBandChartScale()
	testingScaleWithoutKind.Kind = render.ChartScale_UNSPECIFIED_SCALE

	testingScaleWithoutDomain := testingBandChartScale()
	testingScaleWithoutDomain.Domain = nil

	testingScaleWithoutNumericDomain := testingLinearChartScale()
	testingScaleWithoutNumericDomain.Domain = &render.ChartScale_DomainCategories{
		DomainCategories: &render.DomainCategories{Categories: []string{"a1", "a2"}},
	}

	testingScaleWithoutCategoriesDomain := testingBandChartScale()
	testingScaleWithoutCategoriesDomain.Domain = &render.ChartScale_DomainNumeric{
		DomainNumeric: &render.DomainNumeric{
			Start: 20,
			End:   40,
		},
	}

	//nolint: govet
	tt := []struct {
		name              string
		initialChartAxes  *render.ChartAxes
		expectedChartAxes *render.ChartAxes
		expectedErr       error
	}{
		{
			"all_is_set",
			&render.ChartAxes{
				AxisTop:         testingLinearChartScale(),
				AxisTopLabel:    "top",
				AxisBottom:      testingLinearChartScale(),
				AxisBottomLabel: "bottom",
				AxisLeft:        testingBandChartScale(),
				AxisLeftLabel:   "left",
				AxisRight:       testingBandChartScale(),
				AxisRightLabel:  "right",
			},
			&render.ChartAxes{
				AxisTop:         testingLinearChartScaleWithPaddings,
				AxisTopLabel:    "top",
				AxisBottom:      testingLinearChartScaleWithPaddings,
				AxisBottomLabel: "bottom",
				AxisLeft:        testingBandChartScale(),
				AxisLeftLabel:   "left",
				AxisRight:       testingBandChartScale(),
				AxisRightLabel:  "right",
			},
			nil,
		},
		{
			"left_and_bottom_are_set",
			&render.ChartAxes{
				AxisTop:         nil,
				AxisTopLabel:    "",
				AxisBottom:      testingLinearChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        testingBandChartScale(),
				AxisLeftLabel:   "",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			&render.ChartAxes{
				AxisTop:         nil,
				AxisTopLabel:    "",
				AxisBottom:      testingLinearChartScaleWithPaddings,
				AxisBottomLabel: "",
				AxisLeft:        testingBandChartScale(),
				AxisLeftLabel:   "",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			nil,
		},
		{
			"no_axes",
			nil,
			nil,
			apitorenderer.ErrChartAxesAreNotSpecified,
		},
		{
			"no_top_and_bottom",
			&render.ChartAxes{
				AxisTop:         nil,
				AxisTopLabel:    "",
				AxisBottom:      nil,
				AxisBottomLabel: "",
				AxisLeft:        testingLinearChartScale(),
				AxisLeftLabel:   "l",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			nil,
			apitorenderer.ErrChartTopOrBottomAxisShouldBeSpecified,
		},
		{
			"no_left_and_right",
			&render.ChartAxes{
				AxisTop:         testingLinearChartScale(),
				AxisTopLabel:    "t",
				AxisBottom:      nil,
				AxisBottomLabel: "",
				AxisLeft:        nil,
				AxisLeftLabel:   "",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			nil,
			apitorenderer.ErrChartLeftOrRightAxisShouldBeSpecified,
		},
		{
			"no_axis_kind",
			&render.ChartAxes{
				AxisTop:         testingLinearChartScale(),
				AxisTopLabel:    "t",
				AxisBottom:      testingScaleWithoutKind,
				AxisBottomLabel: "b",
				AxisLeft:        testingLinearChartScale(),
				AxisLeftLabel:   "l",
				AxisRight:       testingLinearChartScale(),
				AxisRightLabel:  "r",
			},
			nil,
			apitorenderer.ErrChartAxisKindShouldBeSpecified,
		},
		{
			"top_and_bottom_diff",
			&render.ChartAxes{
				AxisTop:         testingLinearChartScale(),
				AxisTopLabel:    "t",
				AxisBottom:      testingBandChartScale(),
				AxisBottomLabel: "b",
				AxisLeft:        testingLinearChartScale(),
				AxisLeftLabel:   "l",
				AxisRight:       testingLinearChartScale(),
				AxisRightLabel:  "r",
			},
			nil,
			apitorenderer.ErrChartTopAndBottomAxesKindsShouldBeEqual,
		},
		{
			"left_and_right_diff",
			&render.ChartAxes{
				AxisTop:         testingLinearChartScale(),
				AxisTopLabel:    "",
				AxisBottom:      testingLinearChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        testingLinearChartScale(),
				AxisLeftLabel:   "",
				AxisRight:       testingBandChartScale(),
				AxisRightLabel:  "",
			},
			nil,
			apitorenderer.ErrChartLeftAndRightAxesKindsShouldBeEqual,
		},
		{
			"no_domain",
			&render.ChartAxes{
				AxisTop:         testingLinearChartScale(),
				AxisTopLabel:    "",
				AxisBottom:      testingLinearChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        nil,
				AxisLeftLabel:   "",
				AxisRight:       testingScaleWithoutDomain,
				AxisRightLabel:  "",
			},
			nil,
			apitorenderer.ErrChartScaleDomainShouldBeSpecified,
		},
		{
			"no_numeric_domain",
			&render.ChartAxes{
				AxisTop:         testingLinearChartScale(),
				AxisTopLabel:    "",
				AxisBottom:      testingLinearChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        nil,
				AxisLeftLabel:   "",
				AxisRight:       testingScaleWithoutNumericDomain,
				AxisRightLabel:  "",
			},
			nil,
			apitorenderer.ErrChartLinearScaleDomainShouldBeSpecified,
		},
		{
			"no_categories_domain",
			&render.ChartAxes{
				AxisTop:         testingLinearChartScale(),
				AxisTopLabel:    "",
				AxisBottom:      testingLinearChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        nil,
				AxisLeftLabel:   "",
				AxisRight:       testingScaleWithoutCategoriesDomain,
				AxisRightLabel:  "",
			},
			nil,
			apitorenderer.ErrChartBandScaleDomainShouldBeSpecified,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actualChartAxes, actualErr := apitorenderer.ValidateChartAxes(tc.initialChartAxes)
			if tc.expectedChartAxes != nil {
				assert.Equal(t, tc.expectedChartAxes.AxisTop, actualChartAxes.AxisTop)
				assert.Equal(t, tc.expectedChartAxes.AxisBottom, actualChartAxes.AxisBottom)
				assert.Equal(t, tc.expectedChartAxes.AxisLeft, actualChartAxes.AxisLeft)
				assert.Equal(t, tc.expectedChartAxes.AxisRight, actualChartAxes.AxisRight)
			}
			assert.Equal(t, tc.expectedErr, actualErr)
		})
	}
}
