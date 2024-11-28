package controllers

import (
	"net/http"

	"loyalty_system/helpers"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	setResponseHeader(w)

	switch r.Method {
	case http.MethodGet:

		helpers.HealthCheck(w, r)

	default:

		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func setResponseHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}
