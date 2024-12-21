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

func GetUserCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"getUserCategory": "OK",
	}
	SendResponse(w, data, "/v1/user-category/get")
}

func CreateUserCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"createUserCategory": "OK",
	}
	SendResponse(w, data, "/v1/user-category/create")
}

func UpdateUserCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"updateUserCategory": "OK",
	}
	SendResponse(w, data, "/v1/user-category/update")
}

func RemoveUserCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"removeUserCategory": "OK",
	}
	SendResponse(w, data, "/v1/user-category/remove")
}
