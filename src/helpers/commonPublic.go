package helpers

import (
	"encoding/json"
	"io"
	"loyalty_system/logger"
	"loyalty_system/models"
	"net/http"
	"time"
)

func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

func FormatResponse(w http.ResponseWriter, httpStatus int, category string) {
	w.WriteHeader(httpStatus)

	data := ResponseData{
		"error": "Unsuccessfull request",
	}
	SendResponse(w, data, category, httpStatus)
}

func GetCurrentSessionData() map[string]models.User {
	var SessionsMap map[string]map[string]models.User

	resp, err := http.Get("http://authenticator.default.svc.cluster.local:8080/sessions")
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
		logger.Errorf("Error get cookie %v", err)
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
	if ok == false {
		return false, user
	}

	return true, value
}
