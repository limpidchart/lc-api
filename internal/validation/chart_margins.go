package validation

import (
	"errors"
	"fmt"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

const (
	chartMarginMin = 0
	chartMarginMax = 100_000

	marginTopDefault    = 90
	marginBottomDefault = 50
	marginLeftDefault   = 60
	marginRightDefault  = 40
)

var (
	// ErrChartMarginsAreNotSpecified contains error message about not specified chart margins.
	ErrChartMarginsAreNotSpecified = errors.New("chart margins are not specified")

	// ErrChartTopMarginIsTooBig contains error message about too big chart top margin.
	ErrChartTopMarginIsTooBig = fmt.Errorf("chart max top margin is %d", chartMarginMax)

	// ErrChartTopMarginIsTooSmall contains error message about too small chart top margin.
	ErrChartTopMarginIsTooSmall = fmt.Errorf("chart min top margin is %d", chartMarginMin)

	// ErrChartBottomMarginIsTooBig contains error message about too big chart bottom margin.
	ErrChartBottomMarginIsTooBig = fmt.Errorf("chart max bottom margin is %d", chartMarginMax)

	// ErrChartBottomMarginIsTooSmall contains error message about too small chart bottom margin.
	ErrChartBottomMarginIsTooSmall = fmt.Errorf("chart min bottom margin is %d", chartMarginMin)

	// ErrChartLeftMarginIsTooBig contains error message about too big chart left margin.
	ErrChartLeftMarginIsTooBig = fmt.Errorf("chart max left margin is %d", chartMarginMax)

	// ErrChartLeftMarginIsTooSmall contains error message about too small chart left margin.
	ErrChartLeftMarginIsTooSmall = fmt.Errorf("chart min left margin is %d", chartMarginMin)

	// ErrChartRightMarginIsTooBig contains error message about too big chart right margin.
	ErrChartRightMarginIsTooBig = fmt.Errorf("chart max right margin is %d", chartMarginMax)

	// ErrChartRightMarginIsTooSmall contains error message about too small chart right margin.
	ErrChartRightMarginIsTooSmall = fmt.Errorf("chart min right margin is %d", chartMarginMin)
)

// ValidateChartMargins check if every chart margin value is specified and in acceptable range.
func ValidateChartMargins(chartMargins *render.ChartMargins) (*render.ChartMargins, error) {
	if chartMargins == nil {
		return nil, ErrChartMarginsAreNotSpecified
	}

	chartMargins, err := validateTopMargin(chartMargins)
	if err != nil {
		return nil, err
	}

	chartMargins, err = validateBottomMargin(chartMargins)
	if err != nil {
		return nil, err
	}

	chartMargins, err = validateLeftMargin(chartMargins)
	if err != nil {
		return nil, err
	}

	chartMargins, err = validateRightMargin(chartMargins)
	if err != nil {
		return nil, err
	}

	return chartMargins, nil
}

func validateTopMargin(chartMargins *render.ChartMargins) (*render.ChartMargins, error) {
	if chartMargins.MarginTop == nil {
		chartMargins.MarginTop = &wrapperspb.Int32Value{Value: marginTopDefault}

		return chartMargins, nil
	}

	if chartMargins.MarginTop.Value > chartMarginMax {
		return nil, ErrChartTopMarginIsTooBig
	}

	if chartMargins.MarginTop.Value < chartMarginMin {
		return nil, ErrChartTopMarginIsTooSmall
	}

	return chartMargins, nil
}

func validateBottomMargin(chartMargins *render.ChartMargins) (*render.ChartMargins, error) {
	if chartMargins.MarginBottom == nil {
		chartMargins.MarginBottom = &wrapperspb.Int32Value{Value: marginBottomDefault}

		return chartMargins, nil
	}

	if chartMargins.MarginBottom.Value > chartMarginMax {
		return nil, ErrChartBottomMarginIsTooBig
	}

	if chartMargins.MarginBottom.Value < chartMarginMin {
		return nil, ErrChartBottomMarginIsTooSmall
	}

	return chartMargins, nil
}

func validateLeftMargin(chartMargins *render.ChartMargins) (*render.ChartMargins, error) {
	if chartMargins.MarginLeft == nil {
		chartMargins.MarginLeft = &wrapperspb.Int32Value{Value: marginLeftDefault}

		return chartMargins, nil
	}

	if chartMargins.MarginLeft.Value > chartMarginMax {
		return nil, ErrChartLeftMarginIsTooBig
	}

	if chartMargins.MarginLeft.Value < chartMarginMin {
		return nil, ErrChartLeftMarginIsTooSmall
	}

	return chartMargins, nil
}

func validateRightMargin(chartMargins *render.ChartMargins) (*render.ChartMargins, error) {
	if chartMargins.MarginRight == nil {
		chartMargins.MarginRight = &wrapperspb.Int32Value{Value: marginRightDefault}

		return chartMargins, nil
	}

	if chartMargins.MarginRight.Value > chartMarginMax {
		return nil, ErrChartRightMarginIsTooBig
	}

	if chartMargins.MarginRight.Value < chartMarginMin {
		return nil, ErrChartRightMarginIsTooSmall
	}

	return chartMargins, nil
}
