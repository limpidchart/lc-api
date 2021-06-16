package apitorenderer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/testutils"
	"github.com/limpidchart/lc-api/internal/validate/apitorenderer"
)

func TestValidateChartAxes(t *testing.T) {
	t.Parallel()

	//nolint: govet
	tt := []struct {
		name              string
		initialChartAxes  *render.ChartAxes
		expectedChartAxes *render.ChartAxes
		expectedErr       error
	}{
		{
			"all_is_set",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "top",
				AxisBottom:      testutils.LinearChartScale(),
				AxisBottomLabel: "bottom",
				AxisLeft:        testutils.BandChartScale(),
				AxisLeftLabel:   "left",
				AxisRight:       testutils.BandChartScale(),
				AxisRightLabel:  "right",
			},
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScaleWithDefaults(),
				AxisTopLabel:    "top",
				AxisBottom:      testutils.LinearChartScaleWithDefaults(),
				AxisBottomLabel: "bottom",
				AxisLeft:        testutils.BandChartScale(),
				AxisLeftLabel:   "left",
				AxisRight:       testutils.BandChartScale(),
				AxisRightLabel:  "right",
			},
			nil,
		},
		{
			"left_and_bottom_are_set",
			&render.ChartAxes{
				AxisTop:         nil,
				AxisTopLabel:    "",
				AxisBottom:      testutils.LinearChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        testutils.BandChartScale(),
				AxisLeftLabel:   "",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			&render.ChartAxes{
				AxisTop:         nil,
				AxisTopLabel:    "",
				AxisBottom:      testutils.LinearChartScaleWithDefaults(),
				AxisBottomLabel: "",
				AxisLeft:        testutils.BandChartScale(),
				AxisLeftLabel:   "",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			nil,
		},
		{
			"no_axes",
			nil,
			nil,
			apitorenderer.ErrChartAxesAreNotSpecified,
		},
		{
			"no_top_and_bottom",
			&render.ChartAxes{
				AxisTop:         nil,
				AxisTopLabel:    "",
				AxisBottom:      nil,
				AxisBottomLabel: "",
				AxisLeft:        testutils.LinearChartScale(),
				AxisLeftLabel:   "l",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			nil,
			apitorenderer.ErrChartTopOrBottomAxisShouldBeSpecified,
		},
		{
			"no_left_and_right",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "t",
				AxisBottom:      nil,
				AxisBottomLabel: "",
				AxisLeft:        nil,
				AxisLeftLabel:   "",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			nil,
			apitorenderer.ErrChartLeftOrRightAxisShouldBeSpecified,
		},
		{
			"no_axis_kind",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "t",
				AxisBottom:      testutils.BandChartScaleWithoutKind(),
				AxisBottomLabel: "b",
				AxisLeft:        testutils.LinearChartScale(),
				AxisLeftLabel:   "l",
				AxisRight:       testutils.LinearChartScale(),
				AxisRightLabel:  "r",
			},
			nil,
			apitorenderer.ErrChartAxisKindShouldBeSpecified,
		},
		{
			"top_and_bottom_diff",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "t",
				AxisBottom:      testutils.BandChartScale(),
				AxisBottomLabel: "b",
				AxisLeft:        testutils.LinearChartScale(),
				AxisLeftLabel:   "l",
				AxisRight:       testutils.LinearChartScale(),
				AxisRightLabel:  "r",
			},
			nil,
			apitorenderer.ErrChartTopAndBottomAxesKindsShouldBeEqual,
		},
		{
			"left_and_right_diff",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "",
				AxisBottom:      testutils.LinearChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        testutils.LinearChartScale(),
				AxisLeftLabel:   "",
				AxisRight:       testutils.BandChartScale(),
				AxisRightLabel:  "",
			},
			nil,
			apitorenderer.ErrChartLeftAndRightAxesKindsShouldBeEqual,
		},
		{
			"no_domain",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "",
				AxisBottom:      testutils.LinearChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        nil,
				AxisLeftLabel:   "",
				AxisRight:       testutils.BandChartScaleWithoutDomain(),
				AxisRightLabel:  "",
			},
			nil,
			apitorenderer.ErrChartScaleDomainShouldBeSpecified,
		},
		{
			"no_numeric_domain",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "",
				AxisBottom:      testutils.LinearChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        nil,
				AxisLeftLabel:   "",
				AxisRight:       testutils.LinearChartScaleWithoutNumericDomain(),
				AxisRightLabel:  "",
			},
			nil,
			apitorenderer.ErrChartLinearScaleDomainShouldBeSpecified,
		},
		{
			"no_categories_domain",
			&render.ChartAxes{
				AxisTop:         testutils.LinearChartScale(),
				AxisTopLabel:    "",
				AxisBottom:      testutils.LinearChartScale(),
				AxisBottomLabel: "",
				AxisLeft:        nil,
				AxisLeftLabel:   "",
				AxisRight:       testutils.BandChartScaleWithoutCategoriesDomain(),
				AxisRightLabel:  "",
			},
			nil,
			apitorenderer.ErrChartBandScaleDomainShouldBeSpecified,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actualChartAxes, actualErr := apitorenderer.ValidateChartAxes(tc.initialChartAxes)
			if tc.expectedChartAxes != nil {
				assert.Equal(t, tc.expectedChartAxes.AxisTop, actualChartAxes.AxisTop)
				assert.Equal(t, tc.expectedChartAxes.AxisBottom, actualChartAxes.AxisBottom)
				assert.Equal(t, tc.expectedChartAxes.AxisLeft, actualChartAxes.AxisLeft)
				assert.Equal(t, tc.expectedChartAxes.AxisRight, actualChartAxes.AxisRight)
			}
			assert.Equal(t, tc.expectedErr, actualErr)
		})
	}
}
