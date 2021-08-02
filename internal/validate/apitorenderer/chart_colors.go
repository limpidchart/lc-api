package apitorenderer

import (
	"fmt"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/validate/hex"
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

	fillColor, err := validateChartElementColor(chartView.Colors.Fill, defaultColors.Fill)
	if err != nil {
		return nil, err
	}

	strokeColor, err := validateChartElementColor(chartView.Colors.Stroke, defaultColors.Stroke)
	if err != nil {
		return nil, err
	}

	pointFillColor, err := validateChartElementColor(chartView.Colors.PointFill, defaultColors.PointFill)
	if err != nil {
		return nil, err
	}

	pointStrokeColor, err := validateChartElementColor(chartView.Colors.PointStroke, defaultColors.PointStroke)
	if err != nil {
		return nil, err
	}

	chartView.Colors.Fill = fillColor
	chartView.Colors.Stroke = strokeColor
	chartView.Colors.PointFill = pointFillColor
	chartView.Colors.PointStroke = pointStrokeColor

	return chartView, nil
}

func validateChartViewBarsDatasetColors(chartViewBarsDataset *render.ChartViewBarsValues_BarsDataset) (*render.ChartViewBarsValues_BarsDataset, error) {
	if chartViewBarsDataset == nil {
		return nil, nil
	}

	defaultColors := defaultViewColors()

	if chartViewBarsDataset.Colors == nil {
		chartViewBarsDataset.Colors.Fill = defaultColors.Fill
		chartViewBarsDataset.Colors.Stroke = defaultColors.Stroke

		return chartViewBarsDataset, nil
	}

	fillColor, err := validateChartElementColor(chartViewBarsDataset.Colors.Fill, defaultColors.Fill)
	if err != nil {
		return nil, err
	}

	strokeColor, err := validateChartElementColor(chartViewBarsDataset.Colors.Stroke, defaultColors.Stroke)
	if err != nil {
		return nil, err
	}

	chartViewBarsDataset.Colors.Fill = fillColor
	chartViewBarsDataset.Colors.Stroke = strokeColor

	return chartViewBarsDataset, nil
}

func validateChartElementColor(chartElementColor *render.ChartElementColor, defaultColor *render.ChartElementColor) (*render.ChartElementColor, error) {
	if chartElementColor == nil {
		return defaultColor, nil
	}

	if chartElementColor.GetColorHex() == "" && chartElementColor.GetColorRgb() == nil {
		return defaultColor, nil
	}

	if chartElementColor.GetColorHex() != "" {
		validatedChartElementColor, err := hex.ValidateChartElementColor(chartElementColor)
		if err != nil {
			return nil, fmt.Errorf("unable to validate hex chart element color: %w", err)
		}

		return validatedChartElementColor, nil
	}

	if err := rgb.ValidateChartElementColor(chartElementColor); err != nil {
		return nil, fmt.Errorf("unable to validate rgb chart element color: %w", err)
	}

	return chartElementColor, nil
}

func defaultViewColors() *render.ChartViewColors {
	return &render.ChartViewColors{
		Fill: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: fillColorDefault,
			},
		},
		Stroke: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: strokeColorDefault,
			},
		},
		PointFill: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: fillColorDefault,
			},
		},
		PointStroke: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: strokeColorDefault,
			},
		},
	}
}
