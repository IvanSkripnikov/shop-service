package helpers

import (
	"net/http"
)

func GetItemsCategoryList(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"itemsCategoryList": "OK",
	}
	SendResponse(w, data, "/v1/item-category/list", http.StatusOK)
}

func GetItemCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"getItemCategory": "OK",
	}
	SendResponse(w, data, "/v1/item-category/get", http.StatusOK)
}

func CreateItemCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"createItemCategory": "OK",
	}
	SendResponse(w, data, "/v1/item-category/create", http.StatusOK)
}

func UpdateItemCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"updateItemCategory": "OK",
	}
	SendResponse(w, data, "/v1/item-category/update", http.StatusOK)
}

func RemoveItemCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"removeItemCategory": "OK",
	}
	SendResponse(w, data, "/v1/item-category/remove", http.StatusOK)
}
