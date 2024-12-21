package controllers

import (
	"net/http"

	"loyalty_system/helpers"
)

func GetItemsCatogoryListV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetItemsCategoryList(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetItemCategoryV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetItemCategory(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func CreateItemCategoryV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		helpers.CreateItemCategory(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func UpdateItemCategoryV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		helpers.UpdateItemCategory(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func RemoveItemCategoryV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		helpers.RemoveItemCategory(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
