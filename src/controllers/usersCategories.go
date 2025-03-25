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
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/user-category/list")
	}
}

func GetUserCategoryV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetUserCategory(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/user-category/get")
	}
}

func CreateUserCategoryV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		helpers.CreateUserCategory(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/user-category/create")
	}
}

func UpdateUserCategoryV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		helpers.UpdateUserCategory(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/user-category/update")
	}
}

func RemoveUserCategoryV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		helpers.RemoveUserCategory(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/user-category/remove")
	}
}

func GetCategoryByUserV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetCategoryByUser(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/user-category/get-by-user")
	}
}
