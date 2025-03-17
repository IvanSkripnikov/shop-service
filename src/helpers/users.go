package helpers

import (
	"encoding/json"
	"net/http"
	"strings"

	"loyalty_system/models"

	"github.com/IvanSkripnikov/go-gormdb"
	"github.com/IvanSkripnikov/go-logger"
)

func GetUsersList(w http.ResponseWriter, _ *http.Request) {
	category := "/v1/users/list"
	var users []models.User

	db := gormdb.GetClient(models.ServiceDatabase)
	err := db.Find(&users).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": users,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	category := "/v1/users/get"
	var user models.User

	user.ID, _ = getIDFromRequestString(strings.TrimSpace(r.URL.Path))

	db := gormdb.GetClient(models.ServiceDatabase)
	err := db.Where("id = ?", user.ID).First(&user).Error
	if checkError(w, err, category) {
		return
	}

	// проверяем - доступно ли для данного пользователя информация
	auth, authUser := GetAuth(r)
	if auth {
		if authUser.ID != user.ID {
			data := ResponseData{
				"response": nil,
			}
			SendResponse(w, data, category, http.StatusForbidden)
			return
		}
	}

	data := ResponseData{
		"response": user,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func GetMyInfo(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/users/me"

	db := gormdb.GetClient(models.ServiceDatabase)
	err := db.Where("id = ?", user.ID).First(&user).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"data": user,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func UpdateMyInfo(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/users/me/update"

	var userNew models.User
	err := json.NewDecoder(r.Body).Decode(&userNew)

	if checkError(w, err, category) {
		return
	}

	if !isExists("SELECT * FROM users WHERE id = ?", user.ID) {
		FormatResponse(w, http.StatusNotFound, category)
		return
	}

	user = updateUserByNewData(user, userNew)

	currentTimestamp := GetCurrentTimestamp()
	query := "UPDATE users SET username = ?, first_name = ?, last_name = ?, email = ?, phone = ?, updated = ? WHERE id = ?"
	rows, err := DB.Query(query, user.UserName, user.FirstName, user.LastName, user.Email, user.Phone, currentTimestamp, user.ID)

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
	SendResponse(w, data, category, http.StatusOK)
}

func DepositMe(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/users/me/deposit"

	var deposit models.Deposit
	err := json.NewDecoder(r.Body).Decode(&deposit)
	if checkError(w, err, category) {
		return
	}

	// Производим начисление средств через сервис платежей
	err = DepositForAccount(user.ID, deposit.Amount)
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"status": "Success",
	}
	SendResponse(w, data, category, http.StatusOK)
}

func AddLoyalty(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"addLoyalty": "OK",
	}
	SendResponse(w, data, "/v1/users/add-loyalty", http.StatusOK)
}

func RemoveLoyalty(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"removeLoyalty": "OK",
	}
	SendResponse(w, data, "/v1/users/remove-loyalty", http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	category := "/v1/users/create"
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if checkError(w, err, category) {
		return
	}

	// записываем сообщение в БД
	db := gormdb.GetClient(models.ServiceDatabase)
	err = db.Create(&user).Error
	if err != nil {
		logger.Errorf("Cant create new user %v", err)
	}

	data := ResponseData{
		"response": "success",
	}
	SendResponse(w, data, category, http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	category := "/v1/users/update"
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if checkError(w, err, category) {
		return
	}

	if !isExists("SELECT * FROM users WHERE id = ?", user.ID) {
		FormatResponse(w, http.StatusNotFound, category)
		return
	}

	currentTimestamp := GetCurrentTimestamp()
	query := "UPDATE users SET username = ?, first_name = ?, last_name = ?, email = ?, phone = ?, updated = ? WHERE id = ?"
	rows, err := DB.Query(query, user.UserName, user.FirstName, user.LastName, user.Email, user.Phone, currentTimestamp, user.ID)

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
	SendResponse(w, data, category, http.StatusOK)
}

func BlockUser(w http.ResponseWriter, r *http.Request) {
	category := "/v1/users/block"
	var user models.User

	user.ID, _ = getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if user.ID == 0 {
		FormatResponse(w, http.StatusUnprocessableEntity, category)
		return
	}

	if !isExists("SELECT * FROM users WHERE id = ?", user.ID) {
		FormatResponse(w, http.StatusNotFound, category)
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
	SendResponse(w, data, category, http.StatusOK)
}

func ResetPassword(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"resetPassword": "OK",
	}
	SendResponse(w, data, "/v1/users/reset-password", http.StatusOK)
}

func GetStatistics(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"getStatistics": "OK",
	}
	SendResponse(w, data, "/v1/users/statistics", http.StatusOK)
}

func updateUserByNewData(user, userNew models.User) models.User {
	if userNew.Email != "" {
		user.Email = userNew.Email
	}
	if userNew.UserName != "" {
		user.UserName = userNew.UserName
	}
	if userNew.FirstName != "" {
		user.FirstName = userNew.FirstName
	}
	if userNew.LastName != "" {
		user.LastName = userNew.LastName
	}
	if userNew.Phone != "" {
		user.Email = userNew.Email
	}
	if userNew.Phone != "" {
		user.Phone = userNew.Phone
	}

	return user
}
