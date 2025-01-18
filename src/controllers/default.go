package controllers

import (
	"net/http"

	"loyalty_system/helpers"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.HealthCheck(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/health")
	}
}
