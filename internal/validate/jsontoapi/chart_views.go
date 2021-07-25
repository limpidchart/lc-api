package jsontoapi

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
)

const pointValuesCount = 2

var (
	// ErrViewIsEmpty contains error message about empty view.
	ErrViewIsEmpty = errors.New("chart view is empty")

	// ErrValuesShouldBeSpecified contains error message about missing values.
	ErrValuesShouldBeSpecified = errors.New("chart view values are not specified")

	// ErrOnlyOneOfValuesKindShouldBeSpecified contains error message about a case when too many values kinds are provided.
	ErrOnlyOneOfValuesKindShouldBeSpecified = errors.New("only one of bars_values, points_values, scalar_values should be specified")

	// ErrBadPointValuesCount contains error message about bad point.
	ErrBadPointValuesCount = fmt.Errorf("each point should have %d values", pointValuesCount)
)

const (
	viewAreaKind          = "area"
	viewHorizontalBarKind = "horizontal_bar"
	viewLineKind          = "line"
	viewScatterKind       = "scatter"
	viewVerticalBarKind   = "vertical_bar"

	barLabelPositionStartOutside = "start_outside"
	barLabelPositionStartInside  = "start_inside"
	barLabelPositionCenter       = "center"
	barLabelPositionEndInside    = "end_inside"
	barLabelPositionEndOutside   = "end_outside"

	pointTypeCircle = "circle"
	pointTypeSquare = "square"
	pointTypeX      = "x"

	pointLabelPositionTop         = "top"
	pointLabelPositionTopRight    = "top_right"
	pointLabelPositionTopLeft     = "top_left"
	pointLabelPositionLeft        = "left"
	pointLabelPositionRight       = "right"
	pointLabelPositionBottom      = "bottom"
	pointLabelPositionBottomLeft  = "bottom_left"
	pointLabelPositionBottomRight = "bottom_right"

	barLabelVisibleDefault   = true
	pointVisibleDefault      = true
	pointLabelVisibleDefault = true
)

// ChartViewFromJSON parses and validates chart view JSON representation.
func ChartViewFromJSON(jsonView *view.ChartView) (*render.ChartView, error) {
	if jsonView == nil {
		return nil, ErrViewIsEmpty
	}

	var (
		barLabelVisible   bool
		pointVisible      bool
		pointLabelVisible bool
	)

	if jsonView.BarLabelVisible == nil {
		barLabelVisible = barLabelVisibleDefault
	} else {
		barLabelVisible = *jsonView.BarLabelVisible
	}

	if jsonView.PointVisible == nil {
		pointVisible = pointVisibleDefault
	} else {
		pointVisible = *jsonView.PointVisible
	}

	if jsonView.PointLabelVisible == nil {
		pointLabelVisible = pointLabelVisibleDefault
	} else {
		pointLabelVisible = *jsonView.PointLabelVisible
	}

	colors, err := viewColorsFromJSON(jsonView.Colors)
	if err != nil {
		return nil, err
	}

	result := &render.ChartView{
		Kind:               viewKindFromJSON(jsonView.Kind),
		Values:             nil,
		Colors:             colors,
		BarLabelVisible:    &wrapperspb.BoolValue{Value: barLabelVisible},
		BarLabelPosition:   barLabelPositionFromJSON(jsonView.BarLabelPosition),
		PointVisible:       &wrapperspb.BoolValue{Value: pointVisible},
		PointType:          pointTypeFromJSON(jsonView.PointType),
		PointLabelVisible:  &wrapperspb.BoolValue{Value: pointLabelVisible},
		PointLabelPosition: pointLabelPositionFromJSON(jsonView.PointLabelPosition),
	}

	valuesAreSet := false

	if jsonView.BarsValues != nil {
		barsValues, err := barsValuesFromJSON(jsonView)
		if err != nil {
			return nil, err
		}

		result.Values = &render.ChartView_BarsValues{
			BarsValues: barsValues,
		}

		valuesAreSet = true
	}

	if jsonView.PointsValues != nil {
		if valuesAreSet {
			return nil, ErrOnlyOneOfValuesKindShouldBeSpecified
		}

		pointsValues, err := pointValuesFromJSON(jsonView)
		if err != nil {
			return nil, err
		}

		result.Values = &render.ChartView_PointsValues{
			PointsValues: pointsValues,
		}

		valuesAreSet = true
	}

	if jsonView.ScalarValues != nil {
		if valuesAreSet {
			return nil, ErrOnlyOneOfValuesKindShouldBeSpecified
		}

		result.Values = &render.ChartView_ScalarValues{
			ScalarValues: scalarValuesFromJSON(jsonView),
		}

		valuesAreSet = true
	}

	if !valuesAreSet {
		return nil, ErrValuesShouldBeSpecified
	}

	return result, nil
}

func viewKindFromJSON(kind string) render.ChartView_ChartViewKind {
	switch strings.ToLower(kind) {
	case viewAreaKind:
		return render.ChartView_AREA
	case viewHorizontalBarKind:
		return render.ChartView_HORIZONTAL_BAR
	case viewLineKind:
		return render.ChartView_LINE
	case viewScatterKind:
		return render.ChartView_SCATTER
	case viewVerticalBarKind:
		return render.ChartView_VERTICAL_BAR
	default:
		return render.ChartView_UNSPECIFIED_KIND
	}
}

func barLabelPositionFromJSON(position string) render.ChartView_ChartViewBarLabelPosition {
	switch strings.ToLower(position) {
	case barLabelPositionStartOutside:
		return render.ChartView_START_OUTSIDE
	case barLabelPositionStartInside:
		return render.ChartView_START_INSIDE
	case barLabelPositionCenter:
		return render.ChartView_CENTER
	case barLabelPositionEndInside:
		return render.ChartView_END_INSIDE
	case barLabelPositionEndOutside:
		return render.ChartView_END_OUTSIDE
	default:
		return render.ChartView_UNSPECIFIED_BAR_LABEL_POSITION
	}
}

func pointTypeFromJSON(pointType string) render.ChartView_ChartViewPointType {
	switch strings.ToLower(pointType) {
	case pointTypeCircle:
		return render.ChartView_CIRCLE
	case pointTypeSquare:
		return render.ChartView_SQUARE
	case pointTypeX:
		return render.ChartView_X
	default:
		return render.ChartView_UNSPECIFIED_POINT_TYPE
	}
}

func pointLabelPositionFromJSON(position string) render.ChartView_ChartViewPointLabelPosition {
	switch strings.ToLower(position) {
	case pointLabelPositionTop:
		return render.ChartView_TOP
	case pointLabelPositionTopRight:
		return render.ChartView_TOP_RIGHT
	case pointLabelPositionTopLeft:
		return render.ChartView_TOP_LEFT
	case pointLabelPositionLeft:
		return render.ChartView_LEFT
	case pointLabelPositionRight:
		return render.ChartView_RIGHT
	case pointLabelPositionBottom:
		return render.ChartView_BOTTOM
	case pointLabelPositionBottomLeft:
		return render.ChartView_BOTTOM_LEFT
	case pointLabelPositionBottomRight:
		return render.ChartView_BOTTOM_RIGHT
	default:
		return render.ChartView_UNSPECIFIED_POINT_LABEL_POSITION
	}
}

func pointValuesFromJSON(jsonView *view.ChartView) (*render.ChartViewPointsValues, error) {
	points := make([]*render.ChartViewPointsValues_Point, 0, len(jsonView.PointsValues.Values))

	for _, point := range jsonView.PointsValues.Values {
		if len(point) != pointValuesCount {
			return nil, ErrBadPointValuesCount
		}

		points = append(points, &render.ChartViewPointsValues_Point{
			X: point[0],
			Y: point[1],
		})
	}

	return &render.ChartViewPointsValues{
		Points: points,
	}, nil
}

func scalarValuesFromJSON(jsonView *view.ChartView) *render.ChartViewScalarValues {
	return &render.ChartViewScalarValues{
		Values: jsonView.ScalarValues.Values,
	}
}

func barsValuesFromJSON(jsonView *view.ChartView) (*render.ChartViewBarsValues, error) {
	barsDatasets := make([]*render.ChartViewBarsValues_BarsDataset, 0, len(jsonView.BarsValues))

	for _, barsValue := range jsonView.BarsValues {
		fillColor, err := chartElementColorFromJSON(barsValue.FillColor)
		if err != nil {
			return nil, err
		}

		strokeColor, err := chartElementColorFromJSON(barsValue.StrokeColor)
		if err != nil {
			return nil, err
		}

		barsDatasets = append(barsDatasets, &render.ChartViewBarsValues_BarsDataset{
			Values:      barsValue.Values,
			FillColor:   fillColor,
			StrokeColor: strokeColor,
		})
	}

	return &render.ChartViewBarsValues{
		BarsDatasets: barsDatasets,
	}, nil
}
