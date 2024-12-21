package helpers

import (
	"net/http"
)

func GetItemsCategoryList(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"itemsCategoryList": "OK",
	}
	SendResponse(w, data, "/v1/item-category/list")
}

func GetItemCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"getItemCategory": "OK",
	}
	SendResponse(w, data, "/v1/item-category/get")
}

func CreateItemCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"createItemCategory": "OK",
	}
	SendResponse(w, data, "/v1/item-category/create")
}

func UpdateItemCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"updateItemCategory": "OK",
	}
	SendResponse(w, data, "/v1/item-category/update")
}

func RemoveItemCategory(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"removeItemCategory": "OK",
	}
	SendResponse(w, data, "/v1/item-category/remove")
}
