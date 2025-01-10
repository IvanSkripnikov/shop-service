package helpers

import (
	"encoding/json"
	"net/http"
	"strings"

	"loyalty_system/logger"
	"loyalty_system/models"
)

func GetUsersList(w http.ResponseWriter, _ *http.Request) {
	category := "/v1/users/list"
	var users []models.User

	query := "SELECT id, login, password, created, updated, active FROM users WHERE active = 1"
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
		"data": users,
	}
	SendResponse(w, data, category)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	category := "/v1/users/get"
	var user models.User

	user.ID, _ = getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if user.ID == 0 {
		formatResponse(w, http.StatusUnprocessableEntity, category)
		return
	}

	query := "SELECT id, login, created, updated, active from users WHERE id = ? AND active = 1"
	rows, err := DB.Prepare(query)

	if checkError(w, err, category) {
		return
	}

	defer func() {
		_ = rows.Close()
	}()

	err = rows.QueryRow(user.ID).Scan(&user.ID, &user.Login, &user.Created, &user.Updated, &user.Active)
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"data": user,
	}
	SendResponse(w, data, category)
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	category := "/v1/users/create"
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if checkError(w, err, category) {
		return
	}

	query := "INSERT INTO users (login, password, created, updated) VALUES (?, ?, CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP())"
	rows, err := DB.Query(query, user.Login, user.Password)

	if checkError(w, err, category) {
		return
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err()
	}()

	data := ResponseData{
		"message": "User successfully created!",
	}
	SendResponse(w, data, category)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	category := "/v1/users/update"
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if checkError(w, err, category) {
		return
	}

	query := "UPDATE users SET login = ?, updated = CURRENT_TIMESTAMP() WHERE id = ?"
	rows, err := DB.Query(query, user.Login, user.ID)

	if checkError(w, err, category) {
		return
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err()
	}()

	data := ResponseData{
		"message": "User successfully updated!",
	}
	SendResponse(w, data, category)
}

func BlockUser(w http.ResponseWriter, r *http.Request) {
	category := "/v1/users/block"
	var user models.User

	user.ID, _ = getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if user.ID == 0 {
		formatResponse(w, http.StatusUnprocessableEntity, category)
		return
	}

	if !isExists("SELECT * FROM users WHERE id = ?", user.ID) {
		formatResponse(w, http.StatusNotFound, category)
		return
	}

	query := "DELETE FROM users WHERE id = ?"
	rows, err := DB.Query(query, user.ID)

	if checkError(w, err, category) {
		return
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err()
	}()

	data := ResponseData{
		"message": "User successfully removed!",
	}
	SendResponse(w, data, category)
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
