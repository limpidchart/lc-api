package apitorenderer

import (
	"errors"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

const (
	rangeStartDefault = 0

	paddingDefault = 0.1
)

var (
	// ErrChartAxesAreNotSpecified contains error message about not specified chart axes.
	ErrChartAxesAreNotSpecified = errors.New("chart axes are not specified")

	// ErrChartTopOrBottomAxisShouldBeSpecified contains error message about not specified top and bottom chart axes.
	ErrChartTopOrBottomAxisShouldBeSpecified = errors.New("chart top or bottom axis is not specified")

	// ErrChartLeftOrRightAxisShouldBeSpecified contains error message about not specified top and bottom chart axes.
	ErrChartLeftOrRightAxisShouldBeSpecified = errors.New("chart left or right axis is not specified")

	// ErrChartAxisKindShouldBeSpecified contains error message about not specified chart axis kind.
	ErrChartAxisKindShouldBeSpecified = errors.New("chart axis kind is not specified")

	// ErrChartTopAndBottomAxesKindsShouldBeEqual contains error message about different kinds of top and bottom axes.
	ErrChartTopAndBottomAxesKindsShouldBeEqual = errors.New("chart top and bottom axes kind should be equal when both axes are provided")

	// ErrChartLeftAndRightAxesKindsShouldBeEqual contains error message about different kinds of left and right axes.
	ErrChartLeftAndRightAxesKindsShouldBeEqual = errors.New("chart left and right axes kind should be equal when both axes are provided")

	// ErrChartScaleDomainShouldBeSpecified contains error message about not specified domain for scale.
	ErrChartScaleDomainShouldBeSpecified = errors.New("chart scale domain should be specified")

	// ErrChartLinearScaleDomainShouldBeSpecified contains error message about not specified numeric domain for linear scale.
	ErrChartLinearScaleDomainShouldBeSpecified = errors.New("chart linear scale numeric domain should be specified")

	// ErrChartBandScaleDomainShouldBeSpecified contains error message about not specified categories domain for band scale.
	ErrChartBandScaleDomainShouldBeSpecified = errors.New("chart band scale categories domain should be specified")

	// ErrChartSizesAreEmpty contains error message about not specified chart sizes.
	ErrChartSizesAreEmpty = errors.New("chart sizes should be specified to set default ranges")

	// ErrChartMarginsAreEmpty contains error message about not specified chart margins.
	ErrChartMarginsAreEmpty = errors.New("chart margins should be specified to set default ranges")

	// ErrChartWidthIsNotSpecified contains error message about not specified chart width.
	ErrChartWidthIsNotSpecified = errors.New("chart width should be specified to set default ranges for horizontal scale")

	// ErrChartHeightIsNotSpecified contains error message about not specified chart height.
	ErrChartHeightIsNotSpecified = errors.New("chart height should be specified to set default ranges for vertical scale")

	// ErrChartLeftAndRightMarginsAreNotSpecified contains error message about not specified chart left and right margins.
	ErrChartLeftAndRightMarginsAreNotSpecified = errors.New("chart left and right margins should be specified to set default ranges for horizontal scale")

	// ErrChartTopAndBottomMarginsAreNotSpecified contains error message about not specified chart top and bottom margins.
	ErrChartTopAndBottomMarginsAreNotSpecified = errors.New("chart top and bottom margins should be specified to set default ranges for vertical scale")
)

// ValidateChartAxes performs the following:
//  - check if chart axes kinds are valid
//  - checks RangeStart, RangeEnd, InnerPadding, OuterPadding values if they're specified or sets default values
//  - check if scale domain is specified
func ValidateChartAxes(chartAxes *render.ChartAxes, chartSizes *render.ChartSizes, chartMargins *render.ChartMargins) (*render.ChartAxes, error) {
	if chartAxes == nil {
		return nil, ErrChartAxesAreNotSpecified
	}

	if err := validateTopAndBottomAxesKinds(chartAxes); err != nil {
		return nil, err
	}

	if err := validateLeftAndRightAxesKinds(chartAxes); err != nil {
		return nil, err
	}

	if err := validateAxesDomains(chartAxes); err != nil {
		return nil, err
	}

	axisTop, err := setHScaleDefaultRanges(chartAxes.AxisTop, chartSizes, chartMargins)
	if err != nil {
		return nil, err
	}

	axisTop = setChartScaleDefaultPaddings(axisTop)

	axisBottom, err := setHScaleDefaultRanges(chartAxes.AxisBottom, chartSizes, chartMargins)
	if err != nil {
		return nil, err
	}

	axisBottom = setChartScaleDefaultPaddings(axisBottom)

	axisLeft, err := setVScaleDefaultRanges(chartAxes.AxisLeft, chartSizes, chartMargins)
	if err != nil {
		return nil, err
	}

	axisLeft = setChartScaleDefaultPaddings(axisLeft)
	axisLeft = invertVScaleRangesIfNeeded(axisLeft)

	axisRight, err := setVScaleDefaultRanges(chartAxes.AxisRight, chartSizes, chartMargins)
	if err != nil {
		return nil, err
	}

	axisRight = setChartScaleDefaultPaddings(axisRight)
	axisRight = invertVScaleRangesIfNeeded(axisRight)

	chartAxes.AxisTop = axisTop
	chartAxes.AxisBottom = axisBottom
	chartAxes.AxisLeft = axisLeft
	chartAxes.AxisRight = axisRight

	return chartAxes, nil
}

// setHScaleDefaultRanges calculates and sets default ranges for horizontal scale if they're not set.
func setHScaleDefaultRanges(chartScale *render.ChartScale, chartSizes *render.ChartSizes, chartMargins *render.ChartMargins) (*render.ChartScale, error) {
	if chartScale == nil {
		return chartScale, nil
	}

	if chartScale.RangeStart != nil && chartScale.RangeEnd != nil {
		return chartScale, nil
	}

	if chartSizes == nil {
		return nil, ErrChartSizesAreEmpty
	}

	if chartMargins == nil {
		return nil, ErrChartMarginsAreEmpty
	}

	if chartSizes.Width == nil {
		return nil, ErrChartWidthIsNotSpecified
	}

	if chartMargins.MarginLeft == nil || chartMargins.MarginRight == nil {
		return nil, ErrChartLeftAndRightMarginsAreNotSpecified
	}

	chartScale.RangeStart = &wrapperspb.Int32Value{Value: rangeStartDefault}
	chartScale.RangeEnd = &wrapperspb.Int32Value{Value: chartSizes.Width.Value - chartMargins.MarginLeft.Value - chartMargins.MarginRight.Value}

	return chartScale, nil
}

// setVScaleDefaultRanges calculates and sets default ranges for vertical scale if they're not set.
func setVScaleDefaultRanges(chartScale *render.ChartScale, chartSizes *render.ChartSizes, chartMargins *render.ChartMargins) (*render.ChartScale, error) {
	if chartScale == nil {
		return chartScale, nil
	}

	if chartScale.RangeStart != nil && chartScale.RangeEnd != nil {
		return chartScale, nil
	}

	if chartSizes == nil {
		return nil, ErrChartSizesAreEmpty
	}

	if chartMargins == nil {
		return nil, ErrChartMarginsAreEmpty
	}

	if chartSizes.Height == nil {
		return nil, ErrChartHeightIsNotSpecified
	}

	if chartMargins.MarginTop == nil || chartMargins.MarginBottom == nil {
		return nil, ErrChartTopAndBottomMarginsAreNotSpecified
	}

	chartScale.RangeStart = &wrapperspb.Int32Value{Value: rangeStartDefault}
	chartScale.RangeEnd = &wrapperspb.Int32Value{Value: chartSizes.Height.Value - chartMargins.MarginTop.Value - chartMargins.MarginBottom.Value}

	return chartScale, nil
}

// invertVScaleRangesIfNeeded inverts vertical scale ranges if it's a linear scale and range_start < range_end.
// We need it because SVG coordinate system's origin is at left top corner.
func invertVScaleRangesIfNeeded(chartScale *render.ChartScale) *render.ChartScale {
	if chartScale == nil {
		return chartScale
	}

	if chartScale.Kind != render.ChartScale_LINEAR {
		return chartScale
	}

	if chartScale.RangeStart != nil && chartScale.RangeEnd != nil && chartScale.RangeStart.Value < chartScale.RangeEnd.Value {
		chartScale.RangeStart, chartScale.RangeEnd = chartScale.RangeEnd, chartScale.RangeStart
	}

	return chartScale
}

func validateTopAndBottomAxesKinds(chartAxes *render.ChartAxes) error {
	if chartAxes.AxisTop == nil && chartAxes.AxisBottom == nil {
		return ErrChartTopOrBottomAxisShouldBeSpecified
	}

	if err := checkIfScaleKindIsSpecified(chartAxes.AxisTop); err != nil {
		return err
	}

	if err := checkIfScaleKindIsSpecified(chartAxes.AxisBottom); err != nil {
		return err
	}

	if chartAxes.AxisTop != nil && chartAxes.AxisBottom != nil && chartAxes.AxisTop.Kind != chartAxes.AxisBottom.Kind {
		return ErrChartTopAndBottomAxesKindsShouldBeEqual
	}

	return nil
}

func validateLeftAndRightAxesKinds(chartAxes *render.ChartAxes) error {
	if chartAxes.AxisLeft == nil && chartAxes.AxisRight == nil {
		return ErrChartLeftOrRightAxisShouldBeSpecified
	}

	if err := checkIfScaleKindIsSpecified(chartAxes.AxisLeft); err != nil {
		return err
	}

	if err := checkIfScaleKindIsSpecified(chartAxes.AxisRight); err != nil {
		return err
	}

	if chartAxes.AxisLeft != nil && chartAxes.AxisRight != nil && chartAxes.AxisLeft.Kind != chartAxes.AxisRight.Kind {
		return ErrChartLeftAndRightAxesKindsShouldBeEqual
	}

	return nil
}

func checkIfScaleKindIsSpecified(chartScale *render.ChartScale) error {
	if chartScale != nil && chartScale.Kind == render.ChartScale_UNSPECIFIED_SCALE {
		return ErrChartAxisKindShouldBeSpecified
	}

	return nil
}

func validateAxesDomains(chartAxes *render.ChartAxes) error {
	if err := checkScaleDomain(chartAxes.AxisTop); err != nil {
		return err
	}

	if err := checkScaleDomain(chartAxes.AxisBottom); err != nil {
		return err
	}

	if err := checkScaleDomain(chartAxes.AxisLeft); err != nil {
		return err
	}

	return checkScaleDomain(chartAxes.AxisRight)
}

func checkScaleDomain(chartScale *render.ChartScale) error {
	if chartScale == nil {
		return nil
	}

	if chartScale.Domain == nil {
		return ErrChartScaleDomainShouldBeSpecified
	}

	if chartScale.Kind == render.ChartScale_LINEAR && chartScale.GetDomainNumeric() == nil {
		return ErrChartLinearScaleDomainShouldBeSpecified
	}

	if chartScale.Kind == render.ChartScale_BAND {
		if chartScale.GetDomainCategories() == nil {
			return ErrChartBandScaleDomainShouldBeSpecified
		}

		if chartScale.GetDomainCategories() != nil && len(chartScale.GetDomainCategories().Categories) == 0 {
			return ErrChartBandScaleDomainShouldBeSpecified
		}
	}

	return nil
}

func setChartScaleDefaultPaddings(chartScale *render.ChartScale) *render.ChartScale {
	if chartScale == nil {
		return chartScale
	}

	if chartScale.InnerPadding == nil {
		chartScale.InnerPadding = &wrapperspb.FloatValue{Value: paddingDefault}
	}

	if chartScale.OuterPadding == nil {
		chartScale.OuterPadding = &wrapperspb.FloatValue{Value: paddingDefault}
	}

	return chartScale
}
