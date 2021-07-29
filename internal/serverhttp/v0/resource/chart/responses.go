package chart

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
)

// Chart representation.
//
// swagger:response chartRepr
type Chart struct {
	// Chart representation.
	//
	// in: body
	Body struct {
		Chart *view.ChartReply `json:"chart"`
	}
}

// NewCreatedChartFromReply returns a new created chart representation from protobuf reply.
func NewCreatedChartFromReply(chartReply *render.ChartReply) *Chart {
	createdAt := chartReply.CreatedAt.AsTime()
	deletedAt := chartReply.DeletedAt.AsTime()

	chartData := string(chartReply.ChartData)
	if len(chartReply.ChartData) != 0 {
		chartData = base64.StdEncoding.EncodeToString(chartReply.ChartData)
	}

	return &Chart{
		Body: struct {
			Chart *view.ChartReply `json:"chart"`
		}{
			Chart: &view.ChartReply{
				RequestID:   chartReply.RequestId,
				ChartID:     chartReply.ChartId,
				ChartStatus: view.ChartStatusCreated.String(),
				CreatedAt:   &createdAt,
				DeletedAt:   &deletedAt,
				ChartData:   chartData,
			},
		},
	}
}

// MarshalJSON implements the json.Marshaller interface.
func (r *Chart) MarshalJSON() ([]byte, error) {
	res, err := json.Marshal(r.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal chart body into JSON: %w", err)
	}

	return res, nil
}
