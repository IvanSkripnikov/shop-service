package controllers

import (
	"net/http"

	"loyalty_system/helpers"
)

func GetItemsListV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetItemsList(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetItemV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetItem(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func CreateItemV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		helpers.CreateItem(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func UpdateItemV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		helpers.UpdateItem(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func RemoveItemV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		helpers.RemoveItem(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
