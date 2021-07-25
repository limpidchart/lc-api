package convert

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

// RenderChartReplyToAPIChartReply converts RenderChartReply to ChartReply.
func RenderChartReplyToAPIChartReply(reqID, chartID string, ts time.Time, rep *render.RenderChartReply) *render.ChartReply {
	return &render.ChartReply{
		RequestId:   reqID,
		ChartId:     chartID,
		ChartStatus: render.ChartStatus_CREATED,
		CreatedAt:   timestamppb.New(ts),
		DeletedAt:   timestamppb.New(ts), // set deleted at == created until storage backend is implemented
		ChartData:   rep.ChartData,
	}
}
