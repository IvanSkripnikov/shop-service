package helpers

import (
	"net/http"
	"time"
)

func TestError(w http.ResponseWriter, _ *http.Request) {
	FormatResponse(w, http.StatusInternalServerError, "/v1/test/error")
}

func TestLongLatency(w http.ResponseWriter, _ *http.Request) {
	time.After(time.Second)

	data := ResponseData{
		"testLongLatency": "OK",
	}
	SendResponse(w, data, "/v1/test/latency", http.StatusOK)
}
