package helpers

import (
	"net/http"
)

func GetUserCategoryList(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"userCategoryList": "OK",
	}
	SendResponse(w, data, "/v1/user-category/list")
}
