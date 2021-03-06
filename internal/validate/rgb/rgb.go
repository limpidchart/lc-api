package rgb

import (
	"fmt"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
)

const (
	minValue = 0

	maxValue = 255
)

// ErrChartElementColorRGBBadValue contains error message about bad RGB value.
var ErrChartElementColorRGBBadValue = fmt.Errorf("rgb value should be between %d and %d if it's set", minValue, maxValue)

// ValidateChartElementColor validates if chart element color can be used as RGB color.
func ValidateChartElementColor(chartElementColor *render.ChartElementColor) error {
	colorRGB := chartElementColor.GetColorRgb()

	if colorRGB == nil {
		return nil
	}

	if _, err := intColorValue(int(colorRGB.R)); err != nil {
		return err
	}

	if _, err := intColorValue(int(colorRGB.G)); err != nil {
		return err
	}

	if _, err := intColorValue(int(colorRGB.B)); err != nil {
		return err
	}

	return nil
}

// ValidateChartElementColorJSON parses and validates chart element RGB color JSON representation.
func ValidateChartElementColorJSON(color *view.ChartElementColor) (*render.ChartElementColor, error) {
	r, err := intColorValue(color.RGB.R)
	if err != nil {
		return nil, err
	}

	g, err := intColorValue(color.RGB.G)
	if err != nil {
		return nil, err
	}

	b, err := intColorValue(color.RGB.B)
	if err != nil {
		return nil, err
	}

	return &render.ChartElementColor{
		ColorValue: &render.ChartElementColor_ColorRgb{
			ColorRgb: &render.ChartElementColor_RGB{
				R: r,
				G: g,
				B: b,
			},
		},
	}, nil
}

func intColorValue(value int) (uint32, error) {
	if value < minValue || value > maxValue {
		return 0, ErrChartElementColorRGBBadValue
	}

	return uint32(value), nil
}
