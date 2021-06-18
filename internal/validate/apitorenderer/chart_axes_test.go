package apitorenderer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/testutils"
	"github.com/limpidchart/lc-api/internal/validate/apitorenderer"
)

func TestValidateChartAxes(t *testing.T) {
	t.Parallel()

	//nolint: govet
	tt := []struct {
		name              string
		initialChartAxes  *render.ChartAxes
		chartSizes        *render.ChartSizes
		chartMargins      *render.ChartMargins
		expectedChartAxes *render.ChartAxes
		expectedErr       error
	}{
		{
			"all_is_set",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "top",
				AxisBottom:      testutils.LinearChartScale(),
				AxisBottomLabel: "bottom",
				AxisLeft:        testutils.BandChartScale(),
				AxisLeftLabel:   "left",
				AxisRight:       testutils.BandChartScale(),
				AxisRightLabel:  "right",
			},
			nil,
			nil,
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScaleWithDefaults(),
				AxisTopLabel:    "top",
				AxisBottom:      testutils.LinearChartScaleWithDefaults(),
				AxisBottomLabel: "bottom",
				AxisLeft:        testutils.BandChartScale(),
				AxisLeftLabel:   "left",
				AxisRight:       testutils.BandChartScale(),
				AxisRightLabel:  "right",
			},
			nil,
		},
		{
			"left_and_bottom_are_set",
			&render.ChartAxes{
				AxisTop:         nil,
				AxisTopLabel:    "",
				AxisBottom:      testutils.BandChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        testutils.LinearChartScale(),
				AxisLeftLabel:   "",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			nil,
			nil,
			&render.ChartAxes{
				AxisTop:         nil,
				AxisTopLabel:    "",
				AxisBottom:      testutils.BandChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        testutils.LinearChartScaleWithDefaultsAndInvertedRanges(),
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
				AxisLeft:        testutils.LinearChartScale(),
				AxisLeftLabel:   "l",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			nil,
			nil,
			nil,
			apitorenderer.ErrChartTopOrBottomAxisShouldBeSpecified,
		},
		{
			"default_ranges",
			&render.ChartAxes{
				AxisTop:         nil,
				AxisTopLabel:    "",
				AxisBottom:      testutils.BandChartScaleWithoutRanges(),
				AxisBottomLabel: "niz",
				AxisLeft:        testutils.LinearChartScaleWithoutRanges(),
				AxisLeftLabel:   "levo",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			&render.ChartSizes{
				Width:  &wrapperspb.Int32Value{Value: 1200},
				Height: &wrapperspb.Int32Value{Value: 1800},
			},
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 10},
				MarginBottom: &wrapperspb.Int32Value{Value: 20},
				MarginLeft:   &wrapperspb.Int32Value{Value: 25},
				MarginRight:  &wrapperspb.Int32Value{Value: 28},
			},
			&render.ChartAxes{
				AxisTop:         nil,
				AxisTopLabel:    "",
				AxisBottom:      testutils.BandChartScaleWithRanges(0, 1200-25-28),
				AxisBottomLabel: "niz",
				AxisLeft:        testutils.LinearChartScaleWithRangesAndPaddings(1800-10-20, 0),
				AxisLeftLabel:   "levo",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			nil,
		},
		{
			"no_left_and_right",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "t",
				AxisBottom:      nil,
				AxisBottomLabel: "",
				AxisLeft:        nil,
				AxisLeftLabel:   "",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			nil,
			nil,
			nil,
			apitorenderer.ErrChartLeftOrRightAxisShouldBeSpecified,
		},
		{
			"no_axis_kind",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "t",
				AxisBottom:      testutils.BandChartScaleWithoutKind(),
				AxisBottomLabel: "b",
				AxisLeft:        testutils.LinearChartScale(),
				AxisLeftLabel:   "l",
				AxisRight:       testutils.LinearChartScale(),
				AxisRightLabel:  "r",
			},
			nil,
			nil,
			nil,
			apitorenderer.ErrChartAxisKindShouldBeSpecified,
		},
		{
			"top_and_bottom_diff",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "t",
				AxisBottom:      testutils.BandChartScale(),
				AxisBottomLabel: "b",
				AxisLeft:        testutils.LinearChartScale(),
				AxisLeftLabel:   "l",
				AxisRight:       testutils.LinearChartScale(),
				AxisRightLabel:  "r",
			},
			nil,
			nil,
			nil,
			apitorenderer.ErrChartTopAndBottomAxesKindsShouldBeEqual,
		},
		{
			"left_and_right_diff",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "",
				AxisBottom:      testutils.LinearChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        testutils.LinearChartScale(),
				AxisLeftLabel:   "",
				AxisRight:       testutils.BandChartScale(),
				AxisRightLabel:  "",
			},
			nil,
			nil,
			nil,
			apitorenderer.ErrChartLeftAndRightAxesKindsShouldBeEqual,
		},
		{
			"no_domain",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "",
				AxisBottom:      testutils.LinearChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        nil,
				AxisLeftLabel:   "",
				AxisRight:       testutils.BandChartScaleWithoutDomain(),
				AxisRightLabel:  "",
			},
			nil,
			nil,
			nil,
			apitorenderer.ErrChartScaleDomainShouldBeSpecified,
		},
		{
			"no_numeric_domain",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "",
				AxisBottom:      testutils.LinearChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        nil,
				AxisLeftLabel:   "",
				AxisRight:       testutils.LinearChartScaleWithoutNumericDomain(),
				AxisRightLabel:  "",
			},
			nil,
			nil,
			nil,
			apitorenderer.ErrChartLinearScaleDomainShouldBeSpecified,
		},
		{
			"no_categories_domain",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "",
				AxisBottom:      testutils.LinearChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        nil,
				AxisLeftLabel:   "",
				AxisRight:       testutils.BandChartScaleWithoutCategoriesDomain(),
				AxisRightLabel:  "",
			},
			nil,
			nil,
			nil,
			apitorenderer.ErrChartBandScaleDomainShouldBeSpecified,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actualChartAxes, actualErr := apitorenderer.ValidateChartAxes(tc.initialChartAxes, tc.chartSizes, tc.chartMargins)
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
