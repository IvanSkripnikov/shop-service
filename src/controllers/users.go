package controllers

import (
	"net/http"

	"loyalty_system/helpers"
)

func GetMyInfoV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		helpers.GetMyInfo(w, r, user)
	case http.MethodPut:
		helpers.UpdateMyInfo(w, r, user)

	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/users/me")
	}
}

func MyDepositV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodPut:
		helpers.DepositMe(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/users/me/deposit")
	}
}

func GetUsersListV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetUsersList(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/users/list")
	}
}

func GetActiveUsersListV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetActiveUsersList(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/users/get-active")
	}
}

func GetUserV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetUser(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/users/get")
	}
}

func AddLoyaltyV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		helpers.AddLoyalty(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/users/add-loyalty")
	}
}

func RemoveLoyaltyV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		helpers.RemoveLoyalty(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/users/remove-loyalty")
	}
}

func CreateUserV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		helpers.CreateUser(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/users/create")
	}
}

func UpdateUserV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		helpers.UpdateUser(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/users/update")
	}
}

func BlockUserV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		helpers.BlockUser(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/users/block")
	}
}

func ResetUserPasswordV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		helpers.ResetPassword(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/users/reset-password")
	}
}

func GetStatisticsV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetStatistics(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/users/statistics")
	}
}

func UserCategoryUpdateV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		helpers.UserCategoryUpdate(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/users/update")
	}
}
