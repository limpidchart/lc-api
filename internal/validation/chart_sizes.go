package validation

import (
	"fmt"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

const (
	chartSizeMinWidth = 10
	chartSizeMaxWidth = 100_000

	chartSizeMinHeight = 10
	chartSizeMaxHeight = 100_000
)

var (
	// ErrChartSizeWidthTooBig contains error message about too big chart width size.
	ErrChartSizeWidthTooBig = fmt.Errorf("chart size max width is %d", chartSizeMaxWidth)

	// ErrChartSizeWidthTooSmall contains error message about too small chart width size.
	ErrChartSizeWidthTooSmall = fmt.Errorf("chart size min width is %d", chartSizeMinWidth)

	// ErrChartSizeHeightTooBig contains error message about too big chart height size.
	ErrChartSizeHeightTooBig = fmt.Errorf("chart size max height is %d", chartSizeMaxHeight)

	// ErrChartSizeHeightTooSmall contains error message about too small chart height size.
	ErrChartSizeHeightTooSmall = fmt.Errorf("chart size min height is %d", chartSizeMinHeight)
)

// ValidateChartSizes check if every chart size value is in acceptable range.
func ValidateChartSizes(chartSizes *render.ChartSizes) error {
	if chartSizes.Width > chartSizeMaxWidth {
		return ErrChartSizeWidthTooBig
	}

	if chartSizes.Width < chartSizeMinWidth {
		return ErrChartSizeWidthTooSmall
	}

	if chartSizes.Height > chartSizeMaxHeight {
		return ErrChartSizeHeightTooBig
	}

	if chartSizes.Height < chartSizeMinHeight {
		return ErrChartSizeHeightTooSmall
	}

	return nil
}
