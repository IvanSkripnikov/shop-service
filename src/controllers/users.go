package controllers

import (
	"net/http"

	"loyalty_system/helpers"
)

func GetUsersListV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		helpers.GetUsersList(w, r)

	default:

		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
