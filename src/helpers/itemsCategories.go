package helpers

import (
	"encoding/json"
	"loyalty_system/models"
	"net/http"
	"strings"
)

func GetItemsCategoryList(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/item-category/list"
	var itemCategories []models.ItemCategory

	err := GormDB.Where("user_category_id <= ?", user.CategoryID).Find(&itemCategories).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": itemCategories,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func GetItemCategory(w http.ResponseWriter, r *http.Request) {
	category := "/v1/item-category/get"
	var itemCategory models.ItemCategory

	itemCategoryID, _ := getIDFromRequestString(strings.TrimSpace(r.URL.Path))

	err := GormDB.Where("id = ?", itemCategoryID).First(&itemCategory).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": itemCategory,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func CreateItemCategory(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/item-category/create"
	var itemCategory models.ItemCategory

	err := json.NewDecoder(r.Body).Decode(&itemCategory)
	if checkError(w, err, category) {
		return
	}
	itemCategory.Created = GetCurrentDate()
	itemCategory.Updated = GetCurrentDate()

	// проверяем, есть ли права на создание категории товара
	if user.CategoryID == models.UserCategoryStandart {
		FormatResponse(w, http.StatusForbidden, category)
		return
	}

	err = GormDB.Create(&itemCategory).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": models.Success,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func UpdateItemCategory(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/item-category/update"
	var itemCategoryRequest models.ItemCategory

	err := json.NewDecoder(r.Body).Decode(&itemCategoryRequest)
	if checkError(w, err, category) {
		return
	}

	// проверяем, есть ли права на изменение категории товара
	if user.CategoryID == models.UserCategoryStandart {
		FormatResponse(w, http.StatusForbidden, category)
		return
	}

	var itemCategory models.ItemCategory
	err = GormDB.Where("id = ?", itemCategoryRequest.ID).First(&itemCategory).Error
	if checkError(w, err, category) {
		return
	}

	err = GormDB.Model(&itemCategory).Updates(models.ItemCategory{
		Title:          itemCategoryRequest.Title,
		Description:    itemCategoryRequest.Description,
		Active:         itemCategoryRequest.Active,
		UserCategoryID: itemCategory.UserCategoryID,
		Updated:        GetCurrentDate(),
	}).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": models.Success,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func RemoveItemCategory(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/item-category/remove"

	// проверяем, есть ли права на удаление категории товара
	if user.CategoryID == models.UserCategoryStandart {
		FormatResponse(w, http.StatusForbidden, category)
		return
	}

	itemCategoryID, err := getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if checkError(w, err, category) {
		return
	}

	err = GormDB.Delete(&models.ItemCategory{}, itemCategoryID).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": models.Success,
	}
	SendResponse(w, data, category, http.StatusOK)
}
