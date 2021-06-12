package apitorenderer

import (
	"errors"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

const (
	rangeStartDefault = 0
	rangeEndDefault   = 100

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
)

// ValidateChartAxes performs the following:
//  - check if chart axes kinds are valid
//  - checks RangeStart, RangeEnd, InnerPadding, OuterPadding values if they're specified or sets default values
//  - check if scale domain is specified
func ValidateChartAxes(chartAxes *render.ChartAxes) (*render.ChartAxes, error) {
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

	chartAxes.AxisTop = setChartScaleDefaultValues(chartAxes.AxisTop)
	chartAxes.AxisBottom = setChartScaleDefaultValues(chartAxes.AxisBottom)
	chartAxes.AxisLeft = setChartScaleDefaultValues(chartAxes.AxisLeft)
	chartAxes.AxisRight = setChartScaleDefaultValues(chartAxes.AxisRight)

	return chartAxes, nil
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

func setChartScaleDefaultValues(chartScale *render.ChartScale) *render.ChartScale {
	if chartScale == nil {
		return chartScale
	}

	if chartScale.RangeStart == nil {
		chartScale.RangeStart = &wrapperspb.Int32Value{Value: rangeStartDefault}
	}

	if chartScale.RangeEnd == nil {
		chartScale.RangeEnd = &wrapperspb.Int32Value{Value: rangeEndDefault}
	}

	if chartScale.InnerPadding == nil {
		chartScale.InnerPadding = &wrapperspb.FloatValue{Value: paddingDefault}
	}

	if chartScale.OuterPadding == nil {
		chartScale.OuterPadding = &wrapperspb.FloatValue{Value: paddingDefault}
	}

	return chartScale
}
