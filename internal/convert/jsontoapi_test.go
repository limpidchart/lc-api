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
		Title: "Chart",
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
			AxisBottom:      testutils.NewBandChartScale().Unembed(),
			AxisBottomLabel: "Bottom Axis",
			AxisLeft:        testutils.NewLinearChartScale().Unembed(),
			AxisLeftLabel:   "Left Axis",
			AxisRight:       nil,
			AxisRightLabel:  "",
		},
		Views: []*render.ChartView{
			testutils.NewVerticalBarView().SetDefaultPointBools().Unembed(),
			testutils.NewLineView().SetDefaultBarBools().Unembed(),
		},
	}

	actual, err := convert.JSONToCreateChartRequest(
		testutils.NewJSONCreateChartRequest().
			SetTitle().
			SetSizes().
			SetMargins().
			SetBandBottomAxis().
			SetBottomAxisLabel().
			SetLinearLeftAxis().
			SetLeftAxisLabel().
			AddVerticalBarView().
			AddLineView().
			Unembed(),
	)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestJSONToCreateChartRequestEmpty(t *testing.T) {
	t.Parallel()

	actual, err := convert.JSONToCreateChartRequest(nil)
	assert.Equal(t, convert.ErrCreateChartRequestJSONIsEmpty, err)
	assert.Empty(t, actual)
}
