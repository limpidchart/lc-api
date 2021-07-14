package convert_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/convert"
	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/testutils"
)

func TestCreateChartRequestToRenderChartRequest(t *testing.T) {
	t.Parallel()

	expected := &render.RenderChartRequest{
		RequestId: "",
		Title:     "Chart",
		Sizes: &render.ChartSizes{
			Width:  &wrapperspb.Int32Value{Value: 1000},
			Height: &wrapperspb.Int32Value{Value: 800},
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
			AxisLeft:        testutils.NewLinearChartScale().InvertRanges().SetPaddings().Unembed(),
			AxisLeftLabel:   "Left Axis",
			AxisRight:       nil,
			AxisRightLabel:  "",
		},
		Views: []*render.ChartView{
			testutils.NewVerticalBarView().SetDefaultPointParams().SetDefaultColors().Unembed(),
			testutils.NewLineView().SetDefaultBarParams().SetDefaultColors().SetFillAndStrokeColor().Unembed(),
		},
	}

	actual, err := convert.CreateChartRequestToRenderChartRequest(
		testutils.NewCreateChartRequest().
			SetTitle().
			SetSizes().
			SetMargins().
			SetBandBottomAxis().
			SetBottomAxisLabel().
			SetLinearLeftAxis().
			SetLeftAxisLabel().
			AddVerticalBarView().
			AddView(testutils.NewLineView().SetDefaultColors().SetFillAndStrokeColor().Unembed()).
			Unembed(),
	)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
