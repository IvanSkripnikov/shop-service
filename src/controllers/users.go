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

func GetUserV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func AddLoyaltyV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		helpers.AddLoyalty(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func RemoveLoyaltyV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		helpers.RemoveLoyalty(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func CreateUserV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		helpers.CreateUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func UpdateUserV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		helpers.UpdateUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func BlockUserV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		helpers.BlockUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ResetUserPasswordV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		helpers.ResetPassword(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetStatisticsV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetStatistics(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
