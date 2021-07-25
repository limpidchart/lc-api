package rgb_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
	"github.com/limpidchart/lc-api/internal/validate/rgb"
)

func TestChartElementColor(t *testing.T) {
	t.Parallel()

	//nolint: govet
	tt := []struct {
		name              string
		chartElementColor *render.ChartElementColor
		expectedErr       error
	}{
		{
			"good rgb",
			&render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorRgb{
					ColorRgb: &render.ChartElementColor_RGB{
						R: 100,
						G: 125,
						B: 150,
					},
				},
			},
			nil,
		},
		{
			"too_big_r",
			&render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorRgb{
					ColorRgb: &render.ChartElementColor_RGB{
						R: 1000,
						G: 125,
						B: 150,
					},
				},
			},
			rgb.ErrChartElementColorRGBBadValue,
		},
		{
			"too_big_g",
			&render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorRgb{
					ColorRgb: &render.ChartElementColor_RGB{
						R: 100,
						G: 1250,
						B: 150,
					},
				},
			},
			rgb.ErrChartElementColorRGBBadValue,
		},
		{
			"too_big_b",
			&render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorRgb{
					ColorRgb: &render.ChartElementColor_RGB{
						R: 100,
						G: 125,
						B: 1500,
					},
				},
			},
			rgb.ErrChartElementColorRGBBadValue,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actualErr := rgb.ValidateChartElementColor(tc.chartElementColor)
			assert.Equal(t, tc.expectedErr, actualErr)
		})
	}
}

func TestValidateChartElementColorJSON(t *testing.T) {
	t.Parallel()

	//nolint: govet
	tt := []struct {
		name              string
		chartElementColor *view.ChartElementColor
		expectedColor     *render.ChartElementColor
		expectedErr       error
	}{
		{
			"good rgb",
			&view.ChartElementColor{
				Hex: "",
				RGB: &view.RGB{
					R: 10,
					G: 25,
					B: 50,
				},
			},
			&render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorRgb{
					ColorRgb: &render.ChartElementColor_RGB{
						R: 10,
						G: 25,
						B: 50,
					},
				},
			},
			nil,
		},
		{
			"too_small_r",
			&view.ChartElementColor{
				Hex: "",
				RGB: &view.RGB{
					R: -10,
					G: 25,
					B: 50,
				},
			},
			nil,
			rgb.ErrChartElementColorRGBBadValue,
		},
		{
			"too_small_g",
			&view.ChartElementColor{
				Hex: "",
				RGB: &view.RGB{
					R: 10,
					G: -25,
					B: 50,
				},
			},
			nil,
			rgb.ErrChartElementColorRGBBadValue,
		},
		{
			"too_small_b",
			&view.ChartElementColor{
				Hex: "",
				RGB: &view.RGB{
					R: 10,
					G: 25,
					B: -50,
				},
			},
			nil,
			rgb.ErrChartElementColorRGBBadValue,
		},
		{
			"too_big_r",
			&view.ChartElementColor{
				Hex: "",
				RGB: &view.RGB{
					R: 1000,
					G: 25,
					B: 50,
				},
			},
			nil,
			rgb.ErrChartElementColorRGBBadValue,
		},
		{
			"too_big_g",
			&view.ChartElementColor{
				Hex: "",
				RGB: &view.RGB{
					R: 10,
					G: 2500,
					B: 50,
				},
			},
			nil,
			rgb.ErrChartElementColorRGBBadValue,
		},
		{
			"too_big_b",
			&view.ChartElementColor{
				Hex: "",
				RGB: &view.RGB{
					R: 10,
					G: 25,
					B: 5000,
				},
			},
			nil,
			rgb.ErrChartElementColorRGBBadValue,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actualColor, actualErr := rgb.ValidateChartElementColorJSON(tc.chartElementColor)
			assert.Equal(t, tc.expectedColor, actualColor)
			assert.Equal(t, tc.expectedErr, actualErr)
		})
	}
}
