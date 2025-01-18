package helpers

import (
	"net/http"
)

func GetUserCategoryList(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"userCategoryList": "OK",
	}
	SendResponse(w, data, "/v1/user-category/list", http.StatusOK)
}

func GetUserCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"getUserCategory": "OK",
	}
	SendResponse(w, data, "/v1/user-category/get", http.StatusOK)
}

func CreateUserCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"createUserCategory": "OK",
	}
	SendResponse(w, data, "/v1/user-category/create", http.StatusOK)
}

func UpdateUserCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"updateUserCategory": "OK",
	}
	SendResponse(w, data, "/v1/user-category/update", http.StatusOK)
}

func RemoveUserCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"removeUserCategory": "OK",
	}
	SendResponse(w, data, "/v1/user-category/remove", http.StatusOK)
}
