package validation

import (
	"fmt"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

const (
	chartMarginMin = 0
	chartMarginMax = 100_000
)

var (
	// ErrChartTopMarginTooBig contains error message about too big chart top margin.
	ErrChartTopMarginTooBig = fmt.Errorf("chart max top margin is %d", chartMarginMax)

	// ErrChartTopMarginTooSmall contains error message about too small chart top margin.
	ErrChartTopMarginTooSmall = fmt.Errorf("chart min top margin is %d", chartMarginMin)

	// ErrChartBottomMarginTooBig contains error message about too big chart bottom margin.
	ErrChartBottomMarginTooBig = fmt.Errorf("chart max bottom margin is %d", chartMarginMax)

	// ErrChartBottomMarginTooSmall contains error message about too small chart bottom margin.
	ErrChartBottomMarginTooSmall = fmt.Errorf("chart min bottom margin is %d", chartMarginMin)

	// ErrChartLeftMarginTooBig contains error message about too big chart left margin.
	ErrChartLeftMarginTooBig = fmt.Errorf("chart max left margin is %d", chartMarginMax)

	// ErrChartLeftMarginTooSmall contains error message about too small chart left margin.
	ErrChartLeftMarginTooSmall = fmt.Errorf("chart min left margin is %d", chartMarginMin)

	// ErrChartRightMarginTooBig contains error message about too big chart right margin.
	ErrChartRightMarginTooBig = fmt.Errorf("chart max right margin is %d", chartMarginMax)

	// ErrChartRightMarginTooSmall contains error message about too small chart right margin.
	ErrChartRightMarginTooSmall = fmt.Errorf("chart min right margin is %d", chartMarginMin)
)

// ValidateChartMargins check if every chart margin value is in acceptable range.
func ValidateChartMargins(chartMargins *render.ChartMargins) error {
	if chartMargins.MarginTop > chartMarginMax {
		return ErrChartTopMarginTooBig
	}

	if chartMargins.MarginTop < chartMarginMin {
		return ErrChartTopMarginTooSmall
	}

	if chartMargins.MarginBottom > chartMarginMax {
		return ErrChartBottomMarginTooBig
	}

	if chartMargins.MarginBottom < chartMarginMin {
		return ErrChartBottomMarginTooSmall
	}

	if chartMargins.MarginLeft > chartMarginMax {
		return ErrChartLeftMarginTooBig
	}

	if chartMargins.MarginLeft < chartMarginMin {
		return ErrChartLeftMarginTooSmall
	}

	if chartMargins.MarginRight > chartMarginMax {
		return ErrChartRightMarginTooBig
	}

	if chartMargins.MarginRight < chartMarginMin {
		return ErrChartRightMarginTooSmall
	}

	return nil
}
