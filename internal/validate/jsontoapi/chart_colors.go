package jsontoapi

import (
	"errors"
	"fmt"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverrest/view/v0"
	"github.com/limpidchart/lc-api/internal/validate/rgb"
)

// ErrOnlyOneOfHexOrRGBColorShouldBeSpecified contains error message about bad color value.
var ErrOnlyOneOfHexOrRGBColorShouldBeSpecified = errors.New("color can be only one of hex or rgb")

func viewColorsFromJSON(colors *view.ChartViewColors) (*render.ChartViewColors, error) {
	if colors == nil {
		return nil, nil
	}

	fillColor, err := chartElementColorFromJSON(colors.FillColor)
	if err != nil {
		return nil, err
	}

	strokeColor, err := chartElementColorFromJSON(colors.StrokeColor)
	if err != nil {
		return nil, err
	}

	pointFillColor, err := chartElementColorFromJSON(colors.PointFillColor)
	if err != nil {
		return nil, err
	}

	pointStrokeColor, err := chartElementColorFromJSON(colors.PointStrokeColor)
	if err != nil {
		return nil, err
	}

	return &render.ChartViewColors{
		FillColor:        fillColor,
		StrokeColor:      strokeColor,
		PointFillColor:   pointFillColor,
		PointStrokeColor: pointStrokeColor,
	}, nil
}

func chartElementColorFromJSON(color *view.ChartElementColor) (*render.ChartElementColor, error) {
	if color == nil {
		return nil, nil
	}

	if color.Hex == "" && color.RGB == nil {
		return nil, nil
	}

	if color.Hex != "" && color.RGB != nil {
		return nil, ErrOnlyOneOfHexOrRGBColorShouldBeSpecified
	}

	if color.Hex != "" {
		return &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: color.Hex,
			},
		}, nil
	}

	rgbColor, err := rgb.ValidateChartElementColorJSON(color)
	if err != nil {
		return nil, fmt.Errorf("unable to validate rgb chart element color: %w", err)
	}

	return rgbColor, nil
}
