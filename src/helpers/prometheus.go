package helpers

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method"},
	)
	ResponseHttpStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "response_http_status",
			Help: "Total number of HTTP statuses.",
		},
		[]string{"status"},
	)
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: []float64{.5, .95, .99},
		},
		[]string{"method"},
	)
)

func addHttpStatusCodeToPrometheus(httpStatusCode int) {
	ResponseHttpStatus.WithLabelValues(strconv.Itoa(httpStatusCode)).Inc()
}

func RegisterCommonMetrics() {
	prometheus.MustRegister(RequestsTotal, ResponseHttpStatus, RequestDuration)
}
