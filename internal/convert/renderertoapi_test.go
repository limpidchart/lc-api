package convert_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/limpidchart/lc-api/internal/convert"
	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

func TestRenderChartReplyToAPIChartReply(t *testing.T) {
	t.Parallel()

	reqID := uuid.New().String()
	chartID := uuid.New().String()
	now := time.Now().UTC()
	data := []byte("svg data")
	renderChartRep := &render.RenderChartReply{
		RequestId: reqID,
		ChartData: data,
	}

	expected := &render.ChartReply{
		RequestId:   reqID,
		ChartId:     chartID,
		ChartStatus: render.ChartStatus_CREATED,
		CreatedAt:   timestamppb.New(now),
		DeletedAt:   timestamppb.New(now),
		ChartData:   data,
	}

	actual := convert.RenderChartReplyToAPIChartReply(reqID, chartID, now, renderChartRep)
	assert.Equal(t, expected, actual)
}
