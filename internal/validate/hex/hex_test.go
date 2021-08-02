package hex_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
	"github.com/limpidchart/lc-api/internal/validate/hex"
)

func TestValidateChartElementColor(t *testing.T) {
	t.Parallel()

	// nolint: govet
	tt := []struct {
		name                      string
		chartElementColor         *render.ChartElementColor
		expectedChartElementColor *render.ChartElementColor
		expectedErr               error
	}{
		{
			"good_hex",
			&render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorHex{
					ColorHex: "#123ABC",
				},
			},
			&render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorHex{
					ColorHex: "#123abc",
				},
			},
			nil,
		},
		{
			"too_long_hex",
			&render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorHex{
					ColorHex: "#123ABCAAAA",
				},
			},
			nil,
			hex.ErrHexBadValueLen,
		},
		{
			"bad_start_symbol",
			&render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorHex{
					ColorHex: "0000001",
				},
			},
			nil,
			hex.ErrHexDoesntStartWithHash,
		},
		{
			"bad_hex_symbol",
			&render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorHex{
					ColorHex: "#000x00",
				},
			},
			nil,
			hex.ErrHexContainsUnexpectedSymbol,
		},
	}

	for _, tc := range tt {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actualChartElementColor, actualErr := hex.ValidateChartElementColor(tc.chartElementColor)

			assert.Equal(t, tc.expectedErr, actualErr)
			assert.Equal(t, tc.expectedChartElementColor, actualChartElementColor)
		})
	}
}

func TestValidateChartElementColorJSON(t *testing.T) {
	t.Parallel()

	// nolint: govet
	tt := []struct {
		name                      string
		chartElementColor         *view.ChartElementColor
		expectedChartElementColor *render.ChartElementColor
		expectedErr               error
	}{
		{
			"good_hex",
			&view.ChartElementColor{
				Hex: "#123abC",
				RGB: nil,
			},
			&render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorHex{
					ColorHex: "#123abc",
				},
			},
			nil,
		},
		{
			"too_long_hex",
			&view.ChartElementColor{
				Hex: "#123ABCAAAA",
				RGB: nil,
			},
			nil,
			hex.ErrHexBadValueLen,
		},
		{
			"bad_start_symbol",
			&view.ChartElementColor{
				Hex: "0000001",
				RGB: nil,
			},
			nil,
			hex.ErrHexDoesntStartWithHash,
		},
		{
			"bad_hex_symbol",
			&view.ChartElementColor{
				Hex: "#000x00",
				RGB: nil,
			},
			nil,
			hex.ErrHexContainsUnexpectedSymbol,
		},
	}

	for _, tc := range tt {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actualChartElementColor, actualErr := hex.ValidateChartElementColorJSON(tc.chartElementColor)

			assert.Equal(t, tc.expectedErr, actualErr)
			assert.Equal(t, tc.expectedChartElementColor, actualChartElementColor)
		})
	}
}
