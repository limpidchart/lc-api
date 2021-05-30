package validation_test

import (
	"testing"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/validation"
	"github.com/stretchr/testify/assert"
)

func TestValidateChartSizes(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		chartSizes  *render.ChartSizes
		expectedErr error
	}{
		{
			"standard_sizes",
			&render.ChartSizes{
				Width:  800,
				Height: 600,
			},
			nil,
		},
		{
			"small_sizes",
			&render.ChartSizes{
				Width:  10,
				Height: 10,
			},
			nil,
		},
		{
			"big_sizes",
			&render.ChartSizes{
				Width:  100_000,
				Height: 100_000,
			},
			nil,
		},
		{
			"width_is_too_small",
			&render.ChartSizes{
				Width:  -100,
				Height: 100,
			},
			validation.ErrChartSizeWidthTooSmall,
		},
		{
			"width_is_too_big",
			&render.ChartSizes{
				Width:  10_000_000,
				Height: 100,
			},
			validation.ErrChartSizeWidthTooBig,
		},
		{
			"height_is_too_small",
			&render.ChartSizes{
				Width:  100,
				Height: -100,
			},
			validation.ErrChartSizeHeightTooSmall,
		},
		{
			"height_is_too_big",
			&render.ChartSizes{
				Width:  100,
				Height: 10_000_000,
			},
			validation.ErrChartSizeHeightTooBig,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expectedErr, validation.ValidateChartSizes(tc.chartSizes))
		})
	}
}
