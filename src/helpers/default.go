package helpers

import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"status": "success",
	}
	SendResponse(w, data, "/health", http.StatusOK)
}
