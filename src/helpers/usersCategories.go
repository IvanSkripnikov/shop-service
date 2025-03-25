package helpers

import (
	"net/http"
	"strings"

	"loyalty_system/models"
)

func GetUserCategoryList(w http.ResponseWriter, _ *http.Request) {
	category := "/v1/user-category/list"
	var userCategories []models.UserCategory

	err := GormDB.Find(&userCategories).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": userCategories,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func GetUserCategory(w http.ResponseWriter, r *http.Request) {
	category := "/v1/user-category/get"
	var userCategory models.UserCategory

	categoryID, err := getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if checkError(w, err, category) {
		return
	}

	err = GormDB.Where("id = ?", categoryID).First(&userCategory).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": userCategory,
	}
	SendResponse(w, data, category, http.StatusOK)
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

func GetCategoryByUser(w http.ResponseWriter, r *http.Request) {
	category := "/v1/user-category/get-by-user"

	var user models.User
	userID, err := getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if checkError(w, err, category) {
		return
	}
	err = GormDB.Where("id = ?", userID).First(&user).Error
	if checkError(w, err, category) {
		return
	}

	var userCategory models.UserCategory
	err = GormDB.Where("id = ?", user.CategoryID).First(&userCategory).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": userCategory,
	}
	SendResponse(w, data, category, http.StatusOK)
}
