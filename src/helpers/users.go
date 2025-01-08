package helpers

import (
	"loyalty_system/logger"
	"loyalty_system/models"
	"net/http"
)

func GetUsersList(w http.ResponseWriter, _ *http.Request) {
	var users []models.User

	query := "SELECT id, login, password, created, updated, active FROM users"
	rows, err := DB.Query(query)
	if err != nil {
		logger.Error(err.Error())
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err()
	}()

	for rows.Next() {
		user := models.User{}
		if err = rows.Scan(&user.ID, &user.Login, &user.Password, &user.Created, &user.Updated, &user.Active); err != nil {
			logger.Error(err.Error())
			continue
		}
		users = append(users, user)
	}

	data := ResponseData{
		"usersList": users,
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
