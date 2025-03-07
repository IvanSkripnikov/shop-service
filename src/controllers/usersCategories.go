package controllers

import (
	"net/http"

	"loyalty_system/helpers"
)

func GetUserCategoriesListV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetUserCategoryList(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetUserCategoryV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetUserCategory(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func CreateUserCategoryV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		helpers.CreateUserCategory(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func UpdateUserCategoryV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		helpers.UpdateUserCategory(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func RemoveUserCategoryV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		helpers.RemoveUserCategory(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
