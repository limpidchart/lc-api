package jsontoapi

import (
	"errors"
	"fmt"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
	"github.com/limpidchart/lc-api/internal/validate/hex"
	"github.com/limpidchart/lc-api/internal/validate/rgb"
)

// ErrOnlyOneOfHexOrRGBColorShouldBeSpecified contains error message about bad color value.
var ErrOnlyOneOfHexOrRGBColorShouldBeSpecified = errors.New("color can be only one of hex or rgb")

func viewColorsFromJSON(colors *view.ChartViewColors) (*render.ChartViewColors, error) {
	if colors == nil {
		return nil, nil
	}

	fillColor, err := chartElementColorFromJSON(colors.Fill)
	if err != nil {
		return nil, err
	}

	strokeColor, err := chartElementColorFromJSON(colors.Stroke)
	if err != nil {
		return nil, err
	}

	pointFillColor, err := chartElementColorFromJSON(colors.PointFill)
	if err != nil {
		return nil, err
	}

	pointStrokeColor, err := chartElementColorFromJSON(colors.PointStroke)
	if err != nil {
		return nil, err
	}

	return &render.ChartViewColors{
		Fill:        fillColor,
		Stroke:      strokeColor,
		PointFill:   pointFillColor,
		PointStroke: pointStrokeColor,
	}, nil
}

func chartBarsValuesColorsFromJSON(barsValues *view.BarsValues) (*render.ChartViewBarsValues_BarsDataset, error) {
	if barsValues == nil {
		return nil, nil
	}

	if barsValues.Colors == nil {
		return &render.ChartViewBarsValues_BarsDataset{
			Values: barsValues.Values,
			Colors: nil,
		}, nil
	}

	fillColor, err := chartElementColorFromJSON(barsValues.Colors.Fill)
	if err != nil {
		return nil, err
	}

	strokeColor, err := chartElementColorFromJSON(barsValues.Colors.Stroke)
	if err != nil {
		return nil, err
	}

	return &render.ChartViewBarsValues_BarsDataset{
		Values: barsValues.Values,
		Colors: &render.ChartViewBarsValues_ChartViewBarsColors{
			Fill:   fillColor,
			Stroke: strokeColor,
		},
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
		hexColor, err := hex.ValidateChartElementColorJSON(color)
		if err != nil {
			return nil, fmt.Errorf("unable to validate hex chart element color: %w", err)
		}

		return hexColor, nil
	}

	rgbColor, err := rgb.ValidateChartElementColorJSON(color)
	if err != nil {
		return nil, fmt.Errorf("unable to validate rgb chart element color: %w", err)
	}

	return rgbColor, nil
}
