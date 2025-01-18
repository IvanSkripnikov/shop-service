package helpers

import (
	"net/http"

	"loyalty_system/logger"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	responseStatusCode200 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "response_http_status_200",
		Help: "Total number of 200 OK http statuses",
	})
	responseStatusCode400 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "response_http_status_400",
		Help: "Total number of 400 Bad request http statuses",
	})
	responseStatusCode403 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "response_http_status_403",
		Help: "Total number of 403 Forbidden http statuses",
	})
	responseStatusCode404 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "response_http_status_404",
		Help: "Total number of 404 Not found http statuses",
	})
	responseStatusCode405 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "response_http_status_405",
		Help: "Total number of 405 Method Not Allowed http statuses",
	})
	responseStatusCode422 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "response_http_status_422",
		Help: "Total number of 422 Unprocessable Entity http statuses",
	})
	responseStatusCode500 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "response_http_status_500",
		Help: "Total number of 500 Internal Server Error http statuses",
	})
	responseStatusCode502 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "response_http_status_502",
		Help: "Total number of 502 Bad Gateway http statuses",
	})
	responseStatusCodeUnknown = promauto.NewCounter(prometheus.CounterOpts{
		Name: "response_http_status_unknown",
		Help: "Total number of Unknown http statuses",
	})
)

func addHttpStatusCodeToPrometheus(httpStatusCode int) {
	switch httpStatusCode {
	case http.StatusOK:
		responseStatusCode200.Inc()
	case http.StatusBadRequest:
		responseStatusCode400.Inc()
	case http.StatusForbidden:
		responseStatusCode403.Inc()
	case http.StatusNotFound:
		responseStatusCode404.Inc()
	case http.StatusMethodNotAllowed:
		responseStatusCode405.Inc()
	case http.StatusUnprocessableEntity:
		responseStatusCode422.Inc()
	case http.StatusInternalServerError:
		responseStatusCode500.Inc()
	case http.StatusBadGateway:
		responseStatusCode502.Inc()
	default:
		logger.Warningf("Found not tracked http status: %v", httpStatusCode)
		responseStatusCodeUnknown.Inc()
	}

}
