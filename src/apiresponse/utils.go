package apiresponse

import (
	"encoding/json"
	"fmt"
	"net/http"

	"loyalty_system/logger"
)

// InitRoutes Инициализация маршрутов для Rest API.
func InitRoutes(routes map[string]func(http.ResponseWriter, *http.Request)) {
	// Инициализация эндпоинов для Rest API
	for route, handler := range routes {
		http.HandleFunc(route, handler)
	}

	// Вешаем обработчик для главной страницы и логирования ошибочных запросов
	http.HandleFunc("/", homeHandler)
}

// SendResponse Отправить ответ клиенту.
func SendResponse(w http.ResponseWriter, data ResponseData, caption string) {
	response, errEncode := json.Marshal(data)
	if errEncode != nil {
		logger.Error(fmt.Sprintf("Failed to serialize data to get %s. Error: %v", caption, errEncode))
		http.Error(w, errEncode.Error(), http.StatusInternalServerError)

		return
	} else {
		logger.Debug(fmt.Sprintf("Data for receiving %s has been successfully serialized.", caption))
	}

	w.Header().Set("Content-Type", "application/json")
	_, errWrite := w.Write(response)
	if errWrite != nil {
		logger.Error(fmt.Sprintf("Failed to send %s data. Error: %v", caption, errWrite))
		http.Error(w, errWrite.Error(), http.StatusInternalServerError)

		return
	} else {
		logger.Debug(fmt.Sprintf("Data with %s sent successfully.", caption))
	}
}
