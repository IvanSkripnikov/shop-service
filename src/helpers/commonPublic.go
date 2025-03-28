package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"loyalty_system/models"
	"net/http"
	"time"

	logger "github.com/IvanSkripnikov/go-logger"
)

var Config *models.Config

func InitConfig(cfg *models.Config) {
	Config = cfg
}

func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

func GetCurrentDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func FormatResponse(w http.ResponseWriter, httpStatus int, category string) {
	w.WriteHeader(httpStatus)

	data := ResponseData{
		"response": "Unsuccessfull request",
	}
	SendResponse(w, data, category, httpStatus)
}

func GetCurrentSessionData() map[string]models.User {
	var SessionsMap map[string]map[string]models.User

	resp, err := http.Get(Config.RedirectServiceUrl + "/sessions")
	if err != nil {
		logger.Infof("Can't get session data: %v", err)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Infof("Can't parse session data %v", err)
		return nil
	}
	json.Unmarshal(body, &SessionsMap)

	return SessionsMap["sessions"]
}

func GetCurrentSessionID(r *http.Request) string {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		logger.Infof("Cookie session_id not set: %v", err)
		return ""
	}

	return cookie.Value
}

func GetAuth(r *http.Request) (bool, models.User) {
	var user models.User
	sessionID := GetCurrentSessionID(r)
	if sessionID == "" {
		return false, user
	}

	SessionsMap := GetCurrentSessionData()
	value, ok := SessionsMap[sessionID]
	if !ok {
		return false, user
	}

	return true, value
}

func CreateQueryWithScalarResponse(method, url string, data any) (string, error) {
	var err error
	var response string

	jsonData, err := json.Marshal(data)
	if err != nil {
		return response, err
	}
	logger.Infof("json data: %v", string(jsonData))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return response, err
	}

	resp, err := client.Do(req)
	logger.Infof("response for request: %v", resp)
	if err != nil {
		return response, err
	}

	var result map[string]string
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &result)

	logger.Infof("Data from response %v", result)

	// Преобразуем JSON-строку в map
	if err != nil {
		return response, err
	}

	response, ok := result["response"]
	if !ok {
		return "", errors.New("failed to get response")
	}

	return response, nil
}
