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
