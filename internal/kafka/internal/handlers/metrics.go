package handlers

import (
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/api/public"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/presenters"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/services"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/handlers"
	"net/http"
	"strconv"
	"time"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/client/observatorium"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
	"github.com/gorilla/mux"
)

type metricsHandler struct {
	service services.ObservatoriumService
}

func NewMetricsHandler(service services.ObservatoriumService) *metricsHandler {
	return &metricsHandler{
		service: service,
	}
}

func (h metricsHandler) GetMetricsByRangeQuery(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	params := observatorium.MetricsReqParams{}
	query := r.URL.Query()
	cfg := &handlers.HandlerConfig{
		Validate: []handlers.Validate{
			handlers.ValidatQueryParam(query, "duration"),
			handlers.ValidatQueryParam(query, "interval"),
		},
		Action: func() (i interface{}, serviceError *errors.ServiceError) {
			ctx := r.Context()
			params.ResultType = observatorium.RangeQuery
			extractMetricsQueryParams(r, &params)
			kafkaMetrics := &observatorium.KafkaMetrics{}
			foundKafkaId, err := h.service.GetMetricsByKafkaId(ctx, kafkaMetrics, id, params)
			if err != nil {
				return nil, err
			}
			metricList := public.MetricsRangeQueryList{
				Kind: "MetricsRangeQueryList",
				Id:   foundKafkaId,
			}
			metrics, err := presenters.PresentMetricsByRangeQuery(kafkaMetrics)
			if err != nil {
				return nil, err
			}
			metricList.Items = metrics

			return metricList, nil
		},
	}
	handlers.HandleGet(w, r, cfg)
}

func (h metricsHandler) GetMetricsByInstantQuery(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	params := observatorium.MetricsReqParams{}
	cfg := &handlers.HandlerConfig{
		Action: func() (i interface{}, serviceError *errors.ServiceError) {
			ctx := r.Context()
			params.ResultType = observatorium.Query
			extractMetricsQueryParams(r, &params)
			kafkaMetrics := &observatorium.KafkaMetrics{}
			foundKafkaId, err := h.service.GetMetricsByKafkaId(ctx, kafkaMetrics, id, params)
			if err != nil {
				return nil, err
			}
			metricList := public.MetricsInstantQueryList{
				Kind: "MetricsInstantQueryList",
				Id:   foundKafkaId,
			}
			metrics, err := presenters.PresentMetricsByInstantQuery(kafkaMetrics)
			if err != nil {
				return nil, err
			}
			metricList.Items = metrics

			return metricList, nil
		},
	}
	handlers.HandleGet(w, r, cfg)
}

func extractMetricsQueryParams(r *http.Request, q *observatorium.MetricsReqParams) {
	q.FillDefaults()
	queryParams := r.URL.Query()
	if dur := queryParams.Get("duration"); dur != "" {
		if num, err := strconv.ParseInt(dur, 10, 64); err == nil {
			duration := time.Duration(num) * time.Minute
			q.Start = q.End.Add(-duration)
		}
	}
	if step := queryParams.Get("interval"); step != "" {
		if num, err := strconv.Atoi(step); err == nil {
			q.Step = time.Duration(num) * time.Second
		}
	}
	if filters, ok := queryParams["filters"]; ok && len(filters) > 0 {
		q.Filters = filters
	}

}
