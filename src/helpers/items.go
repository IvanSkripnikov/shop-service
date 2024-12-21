package helpers

import (
	"net/http"
)

func GetItemsList(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"itemsList": "OK",
	}
	SendResponse(w, data, "/v1/items/list")
}

func GetItem(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"itemGet": "OK",
	}
	SendResponse(w, data, "/v1/items/get")
}

func CreateItem(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"itemCreate": "OK",
	}
	SendResponse(w, data, "/v1/items/create")
}

func UpdateItem(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"itemUpdate": "OK",
	}
	SendResponse(w, data, "/v1/items/update")
}

func RemoveItem(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"itemRemove": "OK",
	}
	SendResponse(w, data, "/v1/items/remove")
}
