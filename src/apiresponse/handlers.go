package apiresponse

import (
	"fmt"
	"net/http"

	"loyalty_system/logger"
)

// Обработчик для главной страницы и логирования ошибочных запросов.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		logger.Error(fmt.Sprintf("Request for a non-existent page: %s.", r.URL.Path))
		http.NotFound(w, r)

		return
	}

	_, errWrite := w.Write([]byte("Welcome to the Home Page!"))
	if errWrite != nil {
		logger.Error(fmt.Sprintf("Failed to display home page. Error: %v", errWrite))
		http.Error(w, errWrite.Error(), http.StatusInternalServerError)

		return
	} else {
		logger.Debug("Home page displayed successfull.")
	}
}
