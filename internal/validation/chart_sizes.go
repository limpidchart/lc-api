package validation

import (
	"fmt"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

const (
	chartSizeMinWidth = 10
	chartSizeMaxWidth = 100_000

	chartSizeMinHeight = 10
	chartSizeMaxHeight = 100_000

	widthDefault  = 800
	heightDefault = 600
)

var (
	// ErrChartSizeWidthIsTooBig contains error message about too big chart width size.
	ErrChartSizeWidthIsTooBig = fmt.Errorf("chart size max width is %d", chartSizeMaxWidth)

	// ErrChartSizeWidthIsTooSmall contains error message about too small chart width size.
	ErrChartSizeWidthIsTooSmall = fmt.Errorf("chart size min width is %d", chartSizeMinWidth)

	// ErrChartSizeHeightIsTooBig contains error message about too big chart height size.
	ErrChartSizeHeightIsTooBig = fmt.Errorf("chart size max height is %d", chartSizeMaxHeight)

	// ErrChartSizeHeightIsTooSmall contains error message about too small chart height size.
	ErrChartSizeHeightIsTooSmall = fmt.Errorf("chart size min height is %d", chartSizeMinHeight)
)

// ValidateChartSizes check if every chart size value is specified and in acceptable range.
func ValidateChartSizes(chartSizes *render.ChartSizes) (*render.ChartSizes, error) {
	if chartSizes == nil {
		return &render.ChartSizes{
			Width:  &wrapperspb.Int32Value{Value: widthDefault},
			Height: &wrapperspb.Int32Value{Value: heightDefault},
		}, nil
	}

	chartSizes, err := validateChartWidth(chartSizes)
	if err != nil {
		return nil, err
	}

	return validateChartHeight(chartSizes)
}

func validateChartWidth(chartSizes *render.ChartSizes) (*render.ChartSizes, error) {
	if chartSizes.Width == nil {
		chartSizes.Width = &wrapperspb.Int32Value{Value: widthDefault}

		return chartSizes, nil
	}

	if chartSizes.Width.Value > chartSizeMaxWidth {
		return nil, ErrChartSizeWidthIsTooBig
	}

	if chartSizes.Width.Value < chartSizeMinWidth {
		return nil, ErrChartSizeWidthIsTooSmall
	}

	return chartSizes, nil
}

func validateChartHeight(chartSizes *render.ChartSizes) (*render.ChartSizes, error) {
	if chartSizes.Height == nil {
		chartSizes.Height = &wrapperspb.Int32Value{Value: heightDefault}

		return chartSizes, nil
	}

	if chartSizes.Height.Value > chartSizeMaxHeight {
		return nil, ErrChartSizeHeightIsTooBig
	}

	if chartSizes.Height.Value < chartSizeMinHeight {
		return nil, ErrChartSizeHeightIsTooSmall
	}

	return chartSizes, nil
}
