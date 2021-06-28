package convert

import (
	"fmt"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/validate/apitorenderer"
)

// CreateChartRequestToRenderChartRequest validates CreateChartRequest to RenderChartRequest.
func CreateChartRequestToRenderChartRequest(request *render.CreateChartRequest) (*render.RenderChartRequest, error) {
	chartSizes, err := apitorenderer.ValidateChartSizes(request.Sizes)
	if err != nil {
		return nil, fmt.Errorf("unable to validate chart sizes: %w", err)
	}

	chartMargins, err := apitorenderer.ValidateChartMargins(request.Margins)
	if err != nil {
		return nil, fmt.Errorf("unable to validate chart margins: %w", err)
	}

	chartAxes, err := apitorenderer.ValidateChartAxes(request.Axes, request.Sizes, request.Margins)
	if err != nil {
		return nil, fmt.Errorf("unable to validate chart axes: %w", err)
	}

	hScale := selectHorizontalScale(chartAxes)
	vScale := selectVerticalScale(chartAxes)

	categoriesCount := getCategoriesCount(hScale, vScale)

	chartViews, err := apitorenderer.ValidateChartViews(request.Views, categoriesCount, hScale.Kind, vScale.Kind)
	if err != nil {
		return nil, fmt.Errorf("unable to validate chart views: %w", err)
	}

	return &render.RenderChartRequest{
		RequestId: "",
		Title:     request.Title,
		Sizes:     chartSizes,
		Margins:   chartMargins,
		Axes:      chartAxes,
		Views:     chartViews,
	}, nil
}

func selectHorizontalScale(chartAxes *render.ChartAxes) *render.ChartScale {
	if chartAxes.AxisBottom != nil {
		return chartAxes.AxisBottom
	}

	// Can't be nil since it was validated before.
	return chartAxes.AxisTop
}

func selectVerticalScale(chartAxes *render.ChartAxes) *render.ChartScale {
	if chartAxes.AxisLeft != nil {
		return chartAxes.AxisLeft
	}

	// Can't be nil since it was validated before.
	return chartAxes.AxisRight
}

func getCategoriesCount(hScale, vScale *render.ChartScale) int {
	if hScale.Kind == render.ChartScale_BAND {
		return len(hScale.GetDomainCategories().GetCategories())
	}

	if vScale.Kind == render.ChartScale_BAND {
		return len(vScale.GetDomainCategories().GetCategories())
	}

	return 0
}
