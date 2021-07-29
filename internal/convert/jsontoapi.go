package convert

import (
	"errors"
	"fmt"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
	"github.com/limpidchart/lc-api/internal/validate/jsontoapi"
)

// ErrCreateChartRequestJSONIsEmpty represents error message about empty create chart JSON.
var ErrCreateChartRequestJSONIsEmpty = errors.New("create chart JSON is empty")

// JSONToCreateChartRequest converts JSON view representation to *render.CreateChartRequest.
func JSONToCreateChartRequest(reqJSON *view.CreateChartRequest) (*render.CreateChartRequest, error) {
	if reqJSON == nil {
		return nil, ErrCreateChartRequestJSONIsEmpty
	}

	chartAxes, err := chartAxesFromJSON(reqJSON.Chart.Axes)
	if err != nil {
		return nil, err
	}

	chartViews, err := chartViewsFromJSON(reqJSON.Chart.Views)
	if err != nil {
		return nil, err
	}

	return &render.CreateChartRequest{
		Title:   reqJSON.Chart.Title,
		Sizes:   chartSizesFromJSON(reqJSON.Chart.Sizes),
		Margins: chartMarginsFromJSON(reqJSON.Chart.Margins),
		Axes:    chartAxes,
		Views:   chartViews,
	}, nil
}

func chartSizesFromJSON(sizes *view.ChartSizes) *render.ChartSizes {
	if sizes == nil {
		return nil
	}

	var (
		width  *wrapperspb.Int32Value
		height *wrapperspb.Int32Value
	)

	if sizes.Width != nil {
		width = &wrapperspb.Int32Value{Value: int32(*sizes.Width)}
	}

	if sizes.Height != nil {
		height = &wrapperspb.Int32Value{Value: int32(*sizes.Height)}
	}

	return &render.ChartSizes{
		Width:  width,
		Height: height,
	}
}

func chartMarginsFromJSON(margins *view.ChartMargins) *render.ChartMargins {
	if margins == nil {
		return nil
	}

	var (
		marginTop    *wrapperspb.Int32Value
		marginBottom *wrapperspb.Int32Value
		marginLeft   *wrapperspb.Int32Value
		marginRight  *wrapperspb.Int32Value
	)

	if margins.Top != nil {
		marginTop = &wrapperspb.Int32Value{Value: int32(*margins.Top)}
	}

	if margins.Bottom != nil {
		marginBottom = &wrapperspb.Int32Value{Value: int32(*margins.Bottom)}
	}

	if margins.Left != nil {
		marginLeft = &wrapperspb.Int32Value{Value: int32(*margins.Left)}
	}

	if margins.Right != nil {
		marginRight = &wrapperspb.Int32Value{Value: int32(*margins.Right)}
	}

	return &render.ChartMargins{
		MarginTop:    marginTop,
		MarginBottom: marginBottom,
		MarginLeft:   marginLeft,
		MarginRight:  marginRight,
	}
}

func chartAxesFromJSON(axes *view.ChartAxes) (*render.ChartAxes, error) {
	if axes == nil {
		return nil, nil
	}

	axisTop, err := jsontoapi.ChartScaleFromJSON(axes.Top)
	if err != nil {
		return nil, fmt.Errorf("unable to validate top chart scale: %w", err)
	}

	axisBottom, err := jsontoapi.ChartScaleFromJSON(axes.Bottom)
	if err != nil {
		return nil, fmt.Errorf("unable to validate bottom chart scale: %w", err)
	}

	axisLeft, err := jsontoapi.ChartScaleFromJSON(axes.Left)
	if err != nil {
		return nil, fmt.Errorf("unable to validate left chart scale: %w", err)
	}

	axisRight, err := jsontoapi.ChartScaleFromJSON(axes.Right)
	if err != nil {
		return nil, fmt.Errorf("unable to validate right chart scale: %w", err)
	}

	return &render.ChartAxes{
		AxisTop:         axisTop,
		AxisTopLabel:    axes.TopLabel,
		AxisBottom:      axisBottom,
		AxisBottomLabel: axes.BottomLabel,
		AxisLeft:        axisLeft,
		AxisLeftLabel:   axes.LeftLabel,
		AxisRight:       axisRight,
		AxisRightLabel:  axes.RightLabel,
	}, nil
}

func chartViewsFromJSON(views []*view.ChartView) ([]*render.ChartView, error) {
	res := make([]*render.ChartView, 0, len(views))

	for _, jsonView := range views {
		v, err := jsontoapi.ChartViewFromJSON(jsonView)
		if err != nil {
			return nil, fmt.Errorf("unable to validate chart view: %w", err)
		}

		res = append(res, v)
	}

	return res, nil
}
