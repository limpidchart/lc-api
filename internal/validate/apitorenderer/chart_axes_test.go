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
				AxisTop:         testutils.NewLinearChartScale().Unembed(),
				AxisTopLabel:    "top",
				AxisBottom:      testutils.NewLinearChartScale().Unembed(),
				AxisBottomLabel: "bottom",
				AxisLeft:        testutils.NewBandChartScale().Unembed(),
				AxisLeftLabel:   "left",
				AxisRight:       testutils.NewBandChartScale().Unembed(),
				AxisRightLabel:  "right",
			},
			nil,
			nil,
			&render.ChartAxes{
				AxisTop:         testutils.NewLinearChartScale().SetPaddings().Unembed(),
				AxisTopLabel:    "top",
				AxisBottom:      testutils.NewLinearChartScale().SetPaddings().Unembed(),
				AxisBottomLabel: "bottom",
				AxisLeft:        testutils.NewBandChartScale().Unembed(),
				AxisLeftLabel:   "left",
				AxisRight:       testutils.NewBandChartScale().Unembed(),
				AxisRightLabel:  "right",
			},
			nil,
		},
		{
			"left_and_bottom_are_set",
			&render.ChartAxes{
				AxisTop:         nil,
				AxisTopLabel:    "",
				AxisBottom:      testutils.NewBandChartScale().Unembed(),
				AxisBottomLabel: "",
				AxisLeft:        testutils.NewLinearChartScale().Unembed(),
				AxisLeftLabel:   "",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			nil,
			nil,
			&render.ChartAxes{
				AxisTop:         nil,
				AxisTopLabel:    "",
				AxisBottom:      testutils.NewBandChartScale().Unembed(),
				AxisBottomLabel: "",
				AxisLeft:        testutils.NewLinearChartScale().InvertRanges().SetPaddings().Unembed(),
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
				AxisLeft:        testutils.NewLinearChartScale().Unembed(),
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
				AxisBottom:      testutils.NewBandChartScale().UnsetRanges().Unembed(),
				AxisBottomLabel: "niz",
				AxisLeft:        testutils.NewLinearChartScale().UnsetRanges().Unembed(),
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
				AxisBottom:      testutils.NewBandChartScale().SetRanges(0, 1200-25-28).Unembed(),
				AxisBottomLabel: "niz",
				AxisLeft:        testutils.NewLinearChartScale().SetRanges(1800-10-20, 0).SetPaddings().Unembed(),
				AxisLeftLabel:   "levo",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			nil,
		},
		{
			"no_left_and_right",
			&render.ChartAxes{
				AxisTop:         testutils.NewLinearChartScale().Unembed(),
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
				AxisTop:         testutils.NewLinearChartScale().Unembed(),
				AxisTopLabel:    "t",
				AxisBottom:      testutils.NewBandChartScale().UnsetKind().Unembed(),
				AxisBottomLabel: "b",
				AxisLeft:        testutils.NewLinearChartScale().Unembed(),
				AxisLeftLabel:   "l",
				AxisRight:       testutils.NewLinearChartScale().Unembed(),
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
				AxisTop:         testutils.NewLinearChartScale().Unembed(),
				AxisTopLabel:    "t",
				AxisBottom:      testutils.NewBandChartScale().Unembed(),
				AxisBottomLabel: "b",
				AxisLeft:        testutils.NewLinearChartScale().Unembed(),
				AxisLeftLabel:   "l",
				AxisRight:       testutils.NewLinearChartScale().Unembed(),
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
				AxisTop:         testutils.NewLinearChartScale().Unembed(),
				AxisTopLabel:    "",
				AxisBottom:      testutils.NewLinearChartScale().Unembed(),
				AxisBottomLabel: "",
				AxisLeft:        testutils.NewLinearChartScale().Unembed(),
				AxisLeftLabel:   "",
				AxisRight:       testutils.NewBandChartScale().Unembed(),
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
				AxisTop:         testutils.NewLinearChartScale().Unembed(),
				AxisTopLabel:    "",
				AxisBottom:      testutils.NewLinearChartScale().Unembed(),
				AxisBottomLabel: "",
				AxisLeft:        nil,
				AxisLeftLabel:   "",
				AxisRight:       testutils.NewBandChartScale().UnsetDomain().Unembed(),
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
				AxisTop:         testutils.NewLinearChartScale().Unembed(),
				AxisTopLabel:    "",
				AxisBottom:      testutils.NewLinearChartScale().Unembed(),
				AxisBottomLabel: "",
				AxisLeft:        nil,
				AxisLeftLabel:   "",
				AxisRight:       testutils.NewLinearChartScale().UnsetDomain().SetCategoriesDomain().Unembed(),
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
				AxisTop:         testutils.NewLinearChartScale().Unembed(),
				AxisTopLabel:    "",
				AxisBottom:      testutils.NewLinearChartScale().Unembed(),
				AxisBottomLabel: "",
				AxisLeft:        nil,
				AxisLeftLabel:   "",
				AxisRight:       testutils.NewBandChartScale().UnsetDomain().SetNumericDomain().Unembed(),
				AxisRightLabel:  "",
			},
			nil,
			nil,
			nil,
			apitorenderer.ErrChartBandScaleDomainShouldBeSpecified,
		},
		{
			"too big scale label",
			&render.ChartAxes{
				AxisTop:         testutils.NewLinearChartScale().Unembed(),
				AxisTopLabel:    testutils.RandomString(1025),
				AxisBottom:      testutils.NewLinearChartScale().Unembed(),
				AxisBottomLabel: "b",
				AxisLeft:        testutils.NewBandChartScale().Unembed(),
				AxisLeftLabel:   "l",
				AxisRight:       testutils.NewBandChartScale().Unembed(),
				AxisRightLabel:  "r",
			},
			nil,
			nil,
			nil,
			apitorenderer.ErrLabelMaxLen,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actualChartAxes, actualErr := apitorenderer.ValidateChartAxes(tc.initialChartAxes, tc.chartSizes, tc.chartMargins)
			assert.Equal(t, tc.expectedChartAxes, actualChartAxes)
			assert.Equal(t, tc.expectedErr, actualErr)
		})
	}
}
