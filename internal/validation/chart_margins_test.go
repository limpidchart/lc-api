package validation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/validation"
)

func TestValidateChartMargins(t *testing.T) {
	t.Parallel()

	//nolint: govet
	tt := []struct {
		name         string
		chartMargins *render.ChartMargins
		expectedErr  error
	}{
		{
			"standard_margins",
			&render.ChartMargins{
				MarginTop:    90,
				MarginBottom: 50,
				MarginLeft:   60,
				MarginRight:  40,
			},
			nil,
		},
		{
			"big_margins",
			&render.ChartMargins{
				MarginTop:    100_000,
				MarginBottom: 10_000,
				MarginLeft:   100_000,
				MarginRight:  10_000,
			},
			nil,
		},
		{
			"small_margins",
			&render.ChartMargins{
				MarginTop:    0,
				MarginBottom: 1,
				MarginLeft:   0,
				MarginRight:  2,
			},
			nil,
		},
		{
			"too_big_top_margin",
			&render.ChartMargins{
				MarginTop:    10_000_000,
				MarginBottom: 1,
				MarginLeft:   2,
				MarginRight:  3,
			},
			validation.ErrChartTopMarginTooBig,
		},
		{
			"too_big_bottom_margin",
			&render.ChartMargins{
				MarginTop:    1,
				MarginBottom: 1_000_000,
				MarginLeft:   2,
				MarginRight:  3,
			},
			validation.ErrChartBottomMarginTooBig,
		},
		{
			"too_big_left_margin",
			&render.ChartMargins{
				MarginTop:    1,
				MarginBottom: 2,
				MarginLeft:   100_001,
				MarginRight:  3,
			},
			validation.ErrChartLeftMarginTooBig,
		},
		{
			"too_big_right_margin",
			&render.ChartMargins{
				MarginTop:    1,
				MarginBottom: 2,
				MarginLeft:   3,
				MarginRight:  200_000,
			},
			validation.ErrChartRightMarginTooBig,
		},
		{
			"too_small_top_margin",
			&render.ChartMargins{
				MarginTop:    -1,
				MarginBottom: 1,
				MarginLeft:   2,
				MarginRight:  3,
			},
			validation.ErrChartTopMarginTooSmall,
		},
		{
			"too_small_bottom_margin",
			&render.ChartMargins{
				MarginTop:    1,
				MarginBottom: -2,
				MarginLeft:   2,
				MarginRight:  3,
			},
			validation.ErrChartBottomMarginTooSmall,
		},
		{
			"too_small_left_margin",
			&render.ChartMargins{
				MarginTop:    1,
				MarginBottom: 2,
				MarginLeft:   -10_000,
				MarginRight:  3,
			},
			validation.ErrChartLeftMarginTooSmall,
		},
		{
			"too_small_right_margin",
			&render.ChartMargins{
				MarginTop:    1,
				MarginBottom: 2,
				MarginLeft:   3,
				MarginRight:  -100_000,
			},
			validation.ErrChartRightMarginTooSmall,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expectedErr, validation.ValidateChartMargins(tc.chartMargins))
		})
	}
}
