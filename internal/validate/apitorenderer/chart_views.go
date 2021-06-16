package apitorenderer

import (
	"errors"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

const (
	fillColorDefault   = "#71c7ec"
	strokeColorDefault = "#005073"

	barLabelVisibileDefault = true
	barLabelPositionDefault = render.ChartView_CENTER

	pointLabelVisibileDefault = true
	pointLabelPositionDefault = render.ChartView_TOP

	pointVisibileDefault = true
	pointTypeDefault     = render.ChartView_CIRCLE

	rgbMinValue = 0
	rgbMaxValue = 255
)

var (
	// ErrChartViewsAreNotSpecified contains error message about not specified chart views.
	ErrChartViewsAreNotSpecified = errors.New("chart views are not specified")

	// ErrChartViewKindIsUnknown contains error message about unknown chart view kind.
	ErrChartViewKindIsUnknown = errors.New("chart view kind is unknown")

	// ErrChartViewValuesShouldBeSpecified contains error message about not specified chart view values.
	ErrChartViewValuesShouldBeSpecified = errors.New("chart view values should be specified")

	// ErrChartViewValuesCountShouldBeEqualOrLessOfCategoriesCount contains error message about bad amount of chart view values.
	ErrChartViewValuesCountShouldBeEqualOrLessOfCategoriesCount = errors.New("chart view values count should be equal or less than scale categories count")

	// ErrChartElementColorRGBBadValue contains error message about bad RGB value.
	ErrChartElementColorRGBBadValue = errors.New("chart element color RGB value should be between 0 and 255 if it's set")

	// ErrChartScalesForAreaViewAreBad contains error message about bad scales for area view.
	ErrChartScalesForAreaViewAreBad = errors.New("area chart view needs band horizontal scale and linear vertical scale")

	// ErrChartScalesForHorizontalBarViewAreBad contains error message about bad scales for horizontal bar view.
	ErrChartScalesForHorizontalBarViewAreBad = errors.New("horizontal bar chart view needs linear horizontal scale and band vertical scale")

	// ErrChartScalesForLineViewAreBad contains error message about bad scales for line view.
	ErrChartScalesForLineViewAreBad = errors.New("line chart view needs band horizontal scale and linear vertical scale")

	// ErrChartScalesForScatterViewAreBad contains error message about bad scales for scatter view.
	ErrChartScalesForScatterViewAreBad = errors.New("scatter chart view needs linear horizontal scale and linear vertical scale")

	// ErrChartScalesForVerticalBarViewAreBad contains error message about bad scales for vertical bar view.
	ErrChartScalesForVerticalBarViewAreBad = errors.New("vertical bar chart view needs band horizontal scale and linear vertical scale")
)

// ValidateChartViews validates view kind, colors, labels and values and sets defaults if needed.
func ValidateChartViews(chartViews []*render.ChartView, catCount int, hScaleKind, vScaleKind render.ChartScale_ChartScaleKind) ([]*render.ChartView, error) {
	if len(chartViews) == 0 {
		return nil, ErrChartViewsAreNotSpecified
	}

	validatedViews := make([]*render.ChartView, 0, len(chartViews))

	for _, view := range chartViews {
		validatedView, err := validateView(view, catCount, hScaleKind, vScaleKind)
		if err != nil {
			return nil, err
		}

		validatedViews = append(validatedViews, validatedView)
	}

	return validatedViews, nil
}

func validateView(chartView *render.ChartView, catCount int, hScaleKind, vScaleKind render.ChartScale_ChartScaleKind) (*render.ChartView, error) {
	if chartView.Kind == render.ChartView_UNSPECIFIED_KIND {
		return nil, ErrChartViewKindIsUnknown
	}

	if err := validateScalesKindsForView(chartView, hScaleKind, vScaleKind); err != nil {
		return nil, err
	}

	chartView, err := validateViewValues(chartView, catCount)
	if err != nil {
		return nil, err
	}

	chartView, err = validateViewColors(chartView)
	if err != nil {
		return nil, err
	}

	chartView = setViewDefaultBarLabelVisibility(chartView)
	chartView = setViewDefaultBarLabelPosition(chartView)

	chartView = setViewDefaultPointLabelVisibility(chartView)
	chartView = setViewDefaultPointLabelPosition(chartView)

	chartView = setViewDefaultPointVisibility(chartView)
	chartView = setViewDefaultPointType(chartView)

	return chartView, nil
}

func validateScalesKindsForView(chartView *render.ChartView, hScaleKind, vScaleKind render.ChartScale_ChartScaleKind) error {
	switch chartView.Kind {
	case render.ChartView_UNSPECIFIED_KIND:
		return ErrChartViewKindIsUnknown
	case render.ChartView_AREA:
		if hScaleKind != render.ChartScale_BAND || vScaleKind != render.ChartScale_LINEAR {
			return ErrChartScalesForAreaViewAreBad
		}
	case render.ChartView_HORIZONTAL_BAR:
		if hScaleKind != render.ChartScale_LINEAR || vScaleKind != render.ChartScale_BAND {
			return ErrChartScalesForHorizontalBarViewAreBad
		}
	case render.ChartView_LINE:
		if hScaleKind != render.ChartScale_BAND || vScaleKind != render.ChartScale_LINEAR {
			return ErrChartScalesForLineViewAreBad
		}
	case render.ChartView_SCATTER:
		if hScaleKind != render.ChartScale_LINEAR || vScaleKind != render.ChartScale_LINEAR {
			return ErrChartScalesForScatterViewAreBad
		}
	case render.ChartView_VERTICAL_BAR:
		if hScaleKind != render.ChartScale_BAND || vScaleKind != render.ChartScale_LINEAR {
			return ErrChartScalesForVerticalBarViewAreBad
		}
	}

	return nil
}

func validateViewValues(chartView *render.ChartView, catCount int) (*render.ChartView, error) {
	switch chartView.Kind {
	case render.ChartView_UNSPECIFIED_KIND:
		return nil, ErrChartViewKindIsUnknown
	case render.ChartView_AREA:
		return validateScalarValues(chartView, catCount)
	case render.ChartView_HORIZONTAL_BAR:
		return validateViewBarsValues(chartView)
	case render.ChartView_LINE:
		return validateScalarValues(chartView, catCount)
	case render.ChartView_SCATTER:
		return validatePointsValues(chartView)
	case render.ChartView_VERTICAL_BAR:
		return validateViewBarsValues(chartView)
	}

	return nil, ErrChartViewKindIsUnknown
}

func validateScalarValues(chartView *render.ChartView, catCount int) (*render.ChartView, error) {
	if len(chartView.GetScalarValues().Values) == 0 {
		return nil, ErrChartViewValuesShouldBeSpecified
	}

	if len(chartView.GetScalarValues().Values) > catCount {
		return nil, ErrChartViewValuesCountShouldBeEqualOrLessOfCategoriesCount
	}

	return chartView, nil
}

func validatePointsValues(chartView *render.ChartView) (*render.ChartView, error) {
	if len(chartView.GetPointsValues().Points) == 0 {
		return nil, ErrChartViewValuesShouldBeSpecified
	}

	return chartView, nil
}

func validateViewBarsValues(chartView *render.ChartView) (*render.ChartView, error) {
	if chartView.Values == nil {
		return nil, ErrChartViewValuesShouldBeSpecified
	}

	if len(chartView.GetBarsValues().BarsDatasets) == 0 {
		return nil, ErrChartViewValuesShouldBeSpecified
	}

	fixedBarsDatasets := make([]*render.ChartViewBarsValues_BarsDataset, 0, len(chartView.GetBarsValues().BarsDatasets))

	for _, barsDataset := range chartView.GetBarsValues().BarsDatasets {
		if barsDataset.FillColor == nil {
			barsDataset.FillColor = &render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorHex{
					ColorHex: fillColorDefault,
				},
			}
		}

		if barsDataset.StrokeColor == nil {
			barsDataset.StrokeColor = &render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorHex{
					ColorHex: strokeColorDefault,
				},
			}
		}

		fixedBarsDatasets = append(fixedBarsDatasets, barsDataset)
	}

	chartView.Values = &render.ChartView_BarsValues{
		BarsValues: &render.ChartViewBarsValues{
			BarsDatasets: fixedBarsDatasets,
		},
	}

	return chartView, nil
}

func validateViewColors(chartView *render.ChartView) (*render.ChartView, error) {
	defaultColors := defaultViewColors()

	if chartView.Colors == nil {
		chartView.Colors = defaultColors

		return chartView, nil
	}

	fillColor, err := validateColor(chartView.Colors.FillColor, defaultColors.FillColor)
	if err != nil {
		return nil, err
	}

	strokeColor, err := validateColor(chartView.Colors.StrokeColor, defaultColors.StrokeColor)
	if err != nil {
		return nil, err
	}

	pointFillColor, err := validateColor(chartView.Colors.PointFillColor, defaultColors.PointFillColor)
	if err != nil {
		return nil, err
	}

	pointStrokeColor, err := validateColor(chartView.Colors.PointStrokeColor, defaultColors.PointStrokeColor)
	if err != nil {
		return nil, err
	}

	chartView.Colors.FillColor = fillColor
	chartView.Colors.StrokeColor = strokeColor
	chartView.Colors.PointFillColor = pointFillColor
	chartView.Colors.PointStrokeColor = pointStrokeColor

	return chartView, nil
}

func validateColor(chartElementColor *render.ChartElementColor, defaultColor *render.ChartElementColor) (*render.ChartElementColor, error) {
	if chartElementColor == nil {
		return defaultColor, nil
	}

	if chartElementColor.GetColorHex() != "" {
		return chartElementColor, nil
	}

	return validateRGBColor(chartElementColor)
}

func validateRGBColor(chartElementColor *render.ChartElementColor) (*render.ChartElementColor, error) {
	colorRGB := chartElementColor.GetColorRgb()

	if colorRGB == nil {
		return chartElementColor, nil
	}

	if colorRGB.R < rgbMinValue || colorRGB.R > rgbMaxValue {
		return nil, ErrChartElementColorRGBBadValue
	}

	if colorRGB.G < rgbMinValue || colorRGB.G > rgbMaxValue {
		return nil, ErrChartElementColorRGBBadValue
	}

	if colorRGB.B < rgbMinValue || colorRGB.B > rgbMaxValue {
		return nil, ErrChartElementColorRGBBadValue
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

func setViewDefaultBarLabelVisibility(chartView *render.ChartView) *render.ChartView {
	if chartView.BarLabelVisible == nil {
		chartView.BarLabelVisible = &wrapperspb.BoolValue{Value: barLabelVisibileDefault}
	}

	return chartView
}

func setViewDefaultPointLabelVisibility(chartView *render.ChartView) *render.ChartView {
	if chartView.PointLabelVisible == nil {
		chartView.PointLabelVisible = &wrapperspb.BoolValue{Value: pointLabelVisibileDefault}
	}

	return chartView
}

func setViewDefaultBarLabelPosition(chartView *render.ChartView) *render.ChartView {
	if chartView.BarLabelPosition == render.ChartView_UNSPECIFIED_BAR_LABEL_POSITION {
		chartView.BarLabelPosition = barLabelPositionDefault
	}

	return chartView
}

func setViewDefaultPointLabelPosition(chartView *render.ChartView) *render.ChartView {
	if chartView.PointLabelPosition == render.ChartView_UNSPECIFIED_POINT_LABEL_POSITION {
		chartView.PointLabelPosition = pointLabelPositionDefault
	}

	return chartView
}

func setViewDefaultPointVisibility(chartView *render.ChartView) *render.ChartView {
	if chartView.PointVisible == nil {
		chartView.PointVisible = &wrapperspb.BoolValue{Value: pointVisibileDefault}
	}

	return chartView
}

func setViewDefaultPointType(chartView *render.ChartView) *render.ChartView {
	if chartView.PointType == render.ChartView_UNSPECIFIED_POINT_TYPE {
		chartView.PointType = pointTypeDefault
	}

	return chartView
}