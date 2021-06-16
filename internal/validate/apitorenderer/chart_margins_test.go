package apitorenderer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/validate/apitorenderer"
)

func TestValidateChartMargins(t *testing.T) {
	t.Parallel()

	//nolint: govet
	tt := []struct {
		name                 string
		chartMargins         *render.ChartMargins
		expectedChartMargins *render.ChartMargins
		expectedErr          error
	}{
		{
			"standard_margins",
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 90},
				MarginBottom: &wrapperspb.Int32Value{Value: 50},
				MarginLeft:   &wrapperspb.Int32Value{Value: 60},
				MarginRight:  &wrapperspb.Int32Value{Value: 40},
			},
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 90},
				MarginBottom: &wrapperspb.Int32Value{Value: 50},
				MarginLeft:   &wrapperspb.Int32Value{Value: 60},
				MarginRight:  &wrapperspb.Int32Value{Value: 40},
			},
			nil,
		},
		{
			"big_margins",
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 100_000},
				MarginBottom: &wrapperspb.Int32Value{Value: 10_000},
				MarginLeft:   &wrapperspb.Int32Value{Value: 100_000},
				MarginRight:  &wrapperspb.Int32Value{Value: 10_000},
			},
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 100_000},
				MarginBottom: &wrapperspb.Int32Value{Value: 10_000},
				MarginLeft:   &wrapperspb.Int32Value{Value: 100_000},
				MarginRight:  &wrapperspb.Int32Value{Value: 10_000},
			},
			nil,
		},
		{
			"small_margins",
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 0},
				MarginBottom: &wrapperspb.Int32Value{Value: 1},
				MarginLeft:   &wrapperspb.Int32Value{Value: 0},
				MarginRight:  &wrapperspb.Int32Value{Value: 2},
			},
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 0},
				MarginBottom: &wrapperspb.Int32Value{Value: 1},
				MarginLeft:   &wrapperspb.Int32Value{Value: 0},
				MarginRight:  &wrapperspb.Int32Value{Value: 2},
			},
			nil,
		},
		{
			"too_big_top_margin",
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 10_000_000},
				MarginBottom: &wrapperspb.Int32Value{Value: 1},
				MarginLeft:   &wrapperspb.Int32Value{Value: 2},
				MarginRight:  &wrapperspb.Int32Value{Value: 3},
			},
			nil,
			apitorenderer.ErrChartTopMarginIsTooBig,
		},
		{
			"too_big_bottom_margin",
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 1},
				MarginBottom: &wrapperspb.Int32Value{Value: 1_000_000},
				MarginLeft:   &wrapperspb.Int32Value{Value: 2},
				MarginRight:  &wrapperspb.Int32Value{Value: 3},
			},
			nil,
			apitorenderer.ErrChartBottomMarginIsTooBig,
		},
		{
			"too_big_left_margin",
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 1},
				MarginBottom: &wrapperspb.Int32Value{Value: 2},
				MarginLeft:   &wrapperspb.Int32Value{Value: 100_001},
				MarginRight:  &wrapperspb.Int32Value{Value: 3},
			},
			nil,
			apitorenderer.ErrChartLeftMarginIsTooBig,
		},
		{
			"too_big_right_margin",
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 1},
				MarginBottom: &wrapperspb.Int32Value{Value: 2},
				MarginLeft:   &wrapperspb.Int32Value{Value: 3},
				MarginRight:  &wrapperspb.Int32Value{Value: 200_000},
			},
			nil,
			apitorenderer.ErrChartRightMarginIsTooBig,
		},
		{
			"too_small_top_margin",
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: -1},
				MarginBottom: &wrapperspb.Int32Value{Value: 1},
				MarginLeft:   &wrapperspb.Int32Value{Value: 2},
				MarginRight:  &wrapperspb.Int32Value{Value: 3},
			},
			nil,
			apitorenderer.ErrChartTopMarginIsTooSmall,
		},
		{
			"too_small_bottom_margin",
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 1},
				MarginBottom: &wrapperspb.Int32Value{Value: -2},
				MarginLeft:   &wrapperspb.Int32Value{Value: 2},
				MarginRight:  &wrapperspb.Int32Value{Value: 3},
			},
			nil,
			apitorenderer.ErrChartBottomMarginIsTooSmall,
		},
		{
			"too_small_left_margin",
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 1},
				MarginBottom: &wrapperspb.Int32Value{Value: 2},
				MarginLeft:   &wrapperspb.Int32Value{Value: -10_000},
				MarginRight:  &wrapperspb.Int32Value{Value: 3},
			},
			nil,
			apitorenderer.ErrChartLeftMarginIsTooSmall,
		},
		{
			"too_small_right_margin",
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 1},
				MarginBottom: &wrapperspb.Int32Value{Value: 2},
				MarginLeft:   &wrapperspb.Int32Value{Value: 3},
				MarginRight:  &wrapperspb.Int32Value{Value: -100_000},
			},
			nil,
			apitorenderer.ErrChartRightMarginIsTooSmall,
		},
		{
			"no_margins",
			nil,
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 90},
				MarginBottom: &wrapperspb.Int32Value{Value: 50},
				MarginLeft:   &wrapperspb.Int32Value{Value: 60},
				MarginRight:  &wrapperspb.Int32Value{Value: 40},
			},
			nil,
		},
		{
			"default_top_margin",
			&render.ChartMargins{
				MarginTop:    nil,
				MarginBottom: &wrapperspb.Int32Value{Value: 1},
				MarginLeft:   &wrapperspb.Int32Value{Value: 2},
				MarginRight:  &wrapperspb.Int32Value{Value: 3},
			},
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 90},
				MarginBottom: &wrapperspb.Int32Value{Value: 1},
				MarginLeft:   &wrapperspb.Int32Value{Value: 2},
				MarginRight:  &wrapperspb.Int32Value{Value: 3},
			},
			nil,
		},
		{
			"no_bottom_margin",
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 1},
				MarginBottom: nil,
				MarginLeft:   &wrapperspb.Int32Value{Value: 2},
				MarginRight:  &wrapperspb.Int32Value{Value: 3},
			},
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 1},
				MarginBottom: &wrapperspb.Int32Value{Value: 50},
				MarginLeft:   &wrapperspb.Int32Value{Value: 2},
				MarginRight:  &wrapperspb.Int32Value{Value: 3},
			},
			nil,
		},
		{
			"no_left_margin",
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 1},
				MarginBottom: &wrapperspb.Int32Value{Value: 2},
				MarginLeft:   nil,
				MarginRight:  &wrapperspb.Int32Value{Value: 3},
			},
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 1},
				MarginBottom: &wrapperspb.Int32Value{Value: 2},
				MarginLeft:   &wrapperspb.Int32Value{Value: 60},
				MarginRight:  &wrapperspb.Int32Value{Value: 3},
			},
			nil,
		},
		{
			"no_right_margin",
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 1},
				MarginBottom: &wrapperspb.Int32Value{Value: 2},
				MarginLeft:   &wrapperspb.Int32Value{Value: 3},
				MarginRight:  nil,
			},
			&render.ChartMargins{
				MarginTop:    &wrapperspb.Int32Value{Value: 1},
				MarginBottom: &wrapperspb.Int32Value{Value: 2},
				MarginLeft:   &wrapperspb.Int32Value{Value: 3},
				MarginRight:  &wrapperspb.Int32Value{Value: 40},
			},
			nil,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actualChartMargins, actualErr := apitorenderer.ValidateChartMargins(tc.chartMargins)
			if tc.expectedChartMargins != nil {
				assert.Equal(t, tc.expectedChartMargins, actualChartMargins)
			}
			assert.Equal(t, tc.expectedErr, actualErr)
		})
	}
}
