package apitorenderer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/validate/apitorenderer"
)

func TestValidateChartSizes(t *testing.T) {
	t.Parallel()

	//nolint: govet
	tt := []struct {
		name               string
		chartSizes         *render.ChartSizes
		expectedChartSizes *render.ChartSizes
		expectedErr        error
	}{
		{
			"standard_sizes",
			&render.ChartSizes{
				Width:  &wrapperspb.Int32Value{Value: 800},
				Height: &wrapperspb.Int32Value{Value: 600},
			},
			&render.ChartSizes{
				Width:  &wrapperspb.Int32Value{Value: 800},
				Height: &wrapperspb.Int32Value{Value: 600},
			},
			nil,
		},
		{
			"small_sizes",
			&render.ChartSizes{
				Width:  &wrapperspb.Int32Value{Value: 10},
				Height: &wrapperspb.Int32Value{Value: 10},
			},
			&render.ChartSizes{
				Width:  &wrapperspb.Int32Value{Value: 10},
				Height: &wrapperspb.Int32Value{Value: 10},
			},
			nil,
		},
		{
			"big_sizes",
			&render.ChartSizes{
				Width:  &wrapperspb.Int32Value{Value: 100_000},
				Height: &wrapperspb.Int32Value{Value: 100_000},
			},
			&render.ChartSizes{
				Width:  &wrapperspb.Int32Value{Value: 100_000},
				Height: &wrapperspb.Int32Value{Value: 100_000},
			},
			nil,
		},
		{
			"width_is_too_small",
			&render.ChartSizes{
				Width:  &wrapperspb.Int32Value{Value: -100},
				Height: &wrapperspb.Int32Value{Value: 100},
			},
			nil,
			apitorenderer.ErrChartSizeWidthIsTooSmall,
		},
		{
			"width_is_too_big",
			&render.ChartSizes{
				Width:  &wrapperspb.Int32Value{Value: 10_000_000},
				Height: &wrapperspb.Int32Value{Value: 100},
			},
			nil,
			apitorenderer.ErrChartSizeWidthIsTooBig,
		},
		{
			"height_is_too_small",
			&render.ChartSizes{
				Width:  &wrapperspb.Int32Value{Value: 100},
				Height: &wrapperspb.Int32Value{Value: -100},
			},
			nil,
			apitorenderer.ErrChartSizeHeightIsTooSmall,
		},
		{
			"height_is_too_big",
			&render.ChartSizes{
				Width:  &wrapperspb.Int32Value{Value: 100},
				Height: &wrapperspb.Int32Value{Value: 10_000_000},
			},
			nil,
			apitorenderer.ErrChartSizeHeightIsTooBig,
		},
		{
			"no_width",
			&render.ChartSizes{
				Width:  nil,
				Height: &wrapperspb.Int32Value{Value: 100},
			},
			&render.ChartSizes{
				Width:  &wrapperspb.Int32Value{Value: 800},
				Height: &wrapperspb.Int32Value{Value: 100},
			},
			nil,
		},
		{
			"no_height",
			&render.ChartSizes{
				Width:  &wrapperspb.Int32Value{Value: 100},
				Height: nil,
			},
			&render.ChartSizes{
				Width:  &wrapperspb.Int32Value{Value: 100},
				Height: &wrapperspb.Int32Value{Value: 600},
			},
			nil,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actualChartSizes, actualErr := apitorenderer.ValidateChartSizes(tc.chartSizes)
			if tc.expectedChartSizes != nil {
				assert.Equal(t, tc.expectedChartSizes, actualChartSizes)
			}
			assert.Equal(t, tc.expectedErr, actualErr)
		})
	}
}
