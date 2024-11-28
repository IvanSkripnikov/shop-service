package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"loyalty_system/controllers"
)

func TestHealth(t *testing.T) {
	expected := "{\"status\": \"OK\"}"

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	w := httptest.NewRecorder()

	controllers.HealthCheck(w, req)

	res := w.Result()

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("Expected root message but got %v", string(data))
	}
}
