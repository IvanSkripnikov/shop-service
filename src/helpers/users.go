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

func GetUser(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"user": "OK",
	}
	SendResponse(w, data, "/v1/users/get")
}

func AddLoyalty(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"addLoyalty": "OK",
	}
	SendResponse(w, data, "/v1/users/add-loyalty")
}

func RemoveLoyalty(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"removeLoyalty": "OK",
	}
	SendResponse(w, data, "/v1/users/remove-loyalty")
}

func CreateUser(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"createUser": "OK",
	}
	SendResponse(w, data, "/v1/users/create")
}

func UpdateUser(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"updateUser": "OK",
	}
	SendResponse(w, data, "/v1/users/update")
}

func BlockUser(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"blockUser": "OK",
	}
	SendResponse(w, data, "/v1/users/block")
}

func ResetPassword(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"resetPassword": "OK",
	}
	SendResponse(w, data, "/v1/users/reset-password")
}

func GetStatistics(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"getStatistics": "OK",
	}
	SendResponse(w, data, "/v1/users/statistics")
}
