package apitorenderer

import (
	"fmt"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/validate/rgb"
)

const (
	fillColorDefault   = "#71c7ec"
	strokeColorDefault = "#005073"
)

func validateViewColors(chartView *render.ChartView) (*render.ChartView, error) {
	defaultColors := defaultViewColors()

	if chartView.Colors == nil {
		chartView.Colors = defaultColors

		return chartView, nil
	}

	fillColor, err := validateChartElementColor(chartView.Colors.FillColor, defaultColors.FillColor)
	if err != nil {
		return nil, err
	}

	strokeColor, err := validateChartElementColor(chartView.Colors.StrokeColor, defaultColors.StrokeColor)
	if err != nil {
		return nil, err
	}

	pointFillColor, err := validateChartElementColor(chartView.Colors.PointFillColor, defaultColors.PointFillColor)
	if err != nil {
		return nil, err
	}

	pointStrokeColor, err := validateChartElementColor(chartView.Colors.PointStrokeColor, defaultColors.PointStrokeColor)
	if err != nil {
		return nil, err
	}

	chartView.Colors.FillColor = fillColor
	chartView.Colors.StrokeColor = strokeColor
	chartView.Colors.PointFillColor = pointFillColor
	chartView.Colors.PointStrokeColor = pointStrokeColor

	return chartView, nil
}

func validateChartElementColor(chartElementColor *render.ChartElementColor, defaultColor *render.ChartElementColor) (*render.ChartElementColor, error) {
	if chartElementColor == nil {
		return defaultColor, nil
	}

	if chartElementColor.GetColorHex() == "" && chartElementColor.GetColorRgb() == nil {
		return defaultColor, nil
	}

	if chartElementColor.GetColorHex() != "" {
		return chartElementColor, nil
	}

	if err := rgb.ValidateChartElementColor(chartElementColor); err != nil {
		return nil, fmt.Errorf("unable to validate rgb chart element color: %w", err)
	}

	return chartElementColor, nil
}

func defaultViewColors() *render.ChartViewColors {
	return &render.ChartViewColors{
		FillColor: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: fillColorDefault,
			},
		},
		StrokeColor: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: strokeColorDefault,
			},
		},
		PointFillColor: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: fillColorDefault,
			},
		},
		PointStrokeColor: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: strokeColorDefault,
			},
		},
	}
}
