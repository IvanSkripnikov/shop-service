package controllers

import (
	"net/http"

	"loyalty_system/helpers"
)

func TestError(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.TestError(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/test/error")
	}
}

func TestLongLatency(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.TestLongLatency(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/test/latency")
	}
}
