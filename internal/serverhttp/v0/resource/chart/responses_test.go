package chart_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/resource/chart"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
)

func TestNewCreatedChartFromReply(t *testing.T) {
	t.Parallel()

	reqID := "red_id_1"
	chartID := "chart_id_1"
	data := []byte("svg_chart_data_1")
	ts := time.Date(2021, 07, 22, 16, 58, 56, 0, time.UTC)

	expected := &chart.Chart{
		Body: struct {
			Chart *view.ChartReply `json:"chart"`
		}{
			Chart: &view.ChartReply{
				RequestID:   reqID,
				ChartID:     chartID,
				ChartStatus: view.ChartStatusCreated.String(),
				CreatedAt:   &ts,
				DeletedAt:   &ts,
				ChartData:   "svg_chart_data_1",
			},
		},
	}

	reply := &render.ChartReply{
		RequestId:   reqID,
		ChartId:     chartID,
		ChartStatus: render.ChartStatus_CREATED,
		CreatedAt:   timestamppb.New(ts),
		DeletedAt:   timestamppb.New(ts),
		ChartData:   data,
	}

	assert.Equal(t, expected, chart.NewCreatedChartFromReply(reply))
}

func TestChartMarshalJSON(t *testing.T) {
	t.Parallel()

	ts := time.Date(2021, 2, 4, 8, 16, 32, 64, time.UTC)

	tt := []struct {
		name     string
		chart    *chart.Chart
		expected []byte
	}{
		{
			"created_chart",
			&chart.Chart{
				Body: struct {
					Chart *view.ChartReply `json:"chart"`
				}{
					Chart: &view.ChartReply{
						RequestID:   "req_id_3",
						ChartID:     "chart_id_2",
						ChartStatus: view.ChartStatusCreated.String(),
						CreatedAt:   &ts,
						DeletedAt:   &ts,
						ChartData:   "svg_chart_data_2",
					},
				},
			},
			[]byte(`{"chart":{"request_id":"req_id_3","chart_id":"chart_id_2","chart_status":"CREATED","created_at":"2021-02-04T08:16:32.000000064Z","deleted_at":"2021-02-04T08:16:32.000000064Z","chart_data":"svg_chart_data_2"}}`),
		},
		{
			"failed_chart",
			&chart.Chart{
				Body: struct {
					Chart *view.ChartReply `json:"chart"`
				}{
					Chart: &view.ChartReply{
						RequestID:   "req_id_4",
						ChartID:     "",
						ChartStatus: view.ChartStatusError.String(),
						CreatedAt:   nil,
						DeletedAt:   nil,
						ChartData:   "",
					},
				},
			},
			[]byte(`{"chart":{"request_id":"req_id_4","chart_id":"","chart_status":"ERROR","created_at":null,"deleted_at":null,"chart_data":""}}`),
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual, err := json.Marshal(tc.chart)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
