package helpers

import (
	"net/http"
)

func GetUsersList(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"usersList": "OK",
	}
	SendResponse(w, data, "/v1/users/list")
}
