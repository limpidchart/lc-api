package chart

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/renderer"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/middleware"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
)

// Routes implements HTTP handler for charts requests.
func Routes(log *zerolog.Logger, rendererClient render.ChartRendererClient, rendererReqTimeout time.Duration) http.Handler {
	r := chi.NewRouter().
		With(middleware.Recover(log)).
		With(middleware.SetRequestID(log)).
		With(middleware.RequestLogger(log))

	// swagger:route POST /charts Charts createChart
	//
	// Create a new chart
	//
	// Schemes: http, https
	//
	// Produces:
	//   - application/json
	//
	// Responses:
	//   default: error
	//   201: chartRepr
	r.
		With(middleware.RequireCreateChartParams()).
		Post("/", createChartHandler(log, rendererClient, rendererReqTimeout))

	// swagger:route GET /charts/{chart_id} Charts getChart
	//
	// Get chart by ID
	//
	// Schemes: http, https
	//
	// Produces:
	//   - application/json
	//
	// Responses:
	//   default: error
	//   200: chartRepr
	//   404: notFoundError
	r.
		With(middleware.RequireChartID).
		Get(fmt.Sprintf("/{%s}", view.ParamChartID), getChartHandler(log))

	// swagger:route GET /charts Charts listCharts
	//
	// Get charts list
	//
	// Schemes: http, https
	//
	// Produces:
	//   - application/json
	//
	// Responses:
	//   default: error
	r.Get("/", listChartsHandler())

	return r
}

func createChartHandler(log *zerolog.Logger, rendererClient render.ChartRendererClient, rendererReqTimeout time.Duration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := middleware.GetRequestID(r.Context())
		log := log.With().Str(middleware.RequestIDLogKey, reqID).Logger()

		createChartRequest := middleware.GetCreateChartRequest(r.Context())
		if createChartRequest == nil {
			log.Error().Msg("create chart request is empty after middlewares validation")

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			return
		}

		res, err := renderer.CreateChart(r.Context(), renderer.CreateChartOpts{
			RequestID:      reqID,
			Request:        createChartRequest,
			RendererClient: rendererClient,
			Timeout:        rendererReqTimeout,
		})

		switch {
		case err == nil:
			w.WriteHeader(http.StatusCreated)
			middleware.MarshalJSON(w, NewCreatedChartFromReply(res))
		case errors.Is(err, renderer.ErrGenerateChartIDFailed):
			log.Error().Err(err).Msg(fmt.Sprintf("unable to generate a random UUID for %s", view.ParamChartID))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		case errors.Is(err, renderer.ErrCreateChartRequestCancelled):
			log.Warn().Msg("renderer request timed-out")
			http.Error(w, http.StatusText(http.StatusRequestTimeout), http.StatusRequestTimeout)
		default:
			log.Warn().Err(err).Msg("unable to render a chart")
			http.Error(w, fmt.Sprintf("unable to render a chart: %s", err.Error()), http.StatusBadRequest)
		}
	}
}

func getChartHandler(log *zerolog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := middleware.GetRequestID(r.Context())
		log := log.With().Str(middleware.RequestIDLogKey, reqID).Logger()

		chartID := middleware.GetChartID(r.Context())
		if chartID == "" {
			log.Error().Msg("unable to get chart_id from context")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			return
		}

		// We return 404 until storage backend is implemented.
		w.WriteHeader(http.StatusNotFound)
		middleware.MarshalJSON(w, view.NewNotFoundError("chart", chartID))
	}
}

func listChartsHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// We return error until auth is implemented.
		w.WriteHeader(http.StatusNotImplemented)
		middleware.MarshalJSON(w, view.NewError("Unable to get list of charts since auth is not implemented yet"))
	}
}
