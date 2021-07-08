package convert_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/convert"
	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/testutils"
)

func TestJSONToCreateChartRequest(t *testing.T) {
	t.Parallel()

	expected := &render.CreateChartRequest{
		Title: "Vertical and line chart",
		Sizes: &render.ChartSizes{
			Width:  &wrapperspb.Int32Value{Value: 100},
			Height: &wrapperspb.Int32Value{Value: 200},
		},
		Margins: &render.ChartMargins{
			MarginTop:    &wrapperspb.Int32Value{Value: 10},
			MarginBottom: &wrapperspb.Int32Value{Value: 20},
			MarginLeft:   &wrapperspb.Int32Value{Value: 30},
			MarginRight:  &wrapperspb.Int32Value{Value: 40},
		},
		Axes: &render.ChartAxes{
			AxisTop:         nil,
			AxisTopLabel:    "",
			AxisBottom:      testutils.BandChartScale(),
			AxisBottomLabel: "Categories",
			AxisLeft:        testutils.LinearChartScale(),
			AxisLeftLabel:   "Values",
			AxisRight:       nil,
			AxisRightLabel:  "",
		},
		Views: []*render.ChartView{
			testutils.VerticalBarViewWithBoolDefaultsAndEndInsideLabel(),
			testutils.LineViewWithBoolDefaults(),
		},
	}

	actual, err := convert.JSONToCreateChartRequest(testutils.JSONVerticalBarAndLineCreateChartRequest())
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
