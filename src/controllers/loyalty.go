package controllers

import (
	"net/http"

	"loyalty_system/helpers"
)

func GetLoyaltyListV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		helpers.GetLoyaltyList(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/loyalty/list")
	}
}

func GetLoyaltyV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		helpers.GetLoyalty(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/loyalty/get")
	}
}

func GetLoyaltyForUserV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		helpers.GetLoyaltyForUser(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/loyalty/get-for-user")
	}
}

func CreateLoyaltyV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodPost:
		helpers.CreateLoyalty(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/loyalty/create")
	}
}

func UpdateLoyaltyV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodPut:
		helpers.UpdateLoyalty(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/loyalty/update")
	}
}

func RemoveLoyaltyV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodDelete:
		helpers.RemoveLoyalty(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/loyalty/remove")
	}
}

func GetLoyaltyConfigurationListV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		helpers.GetLoyaltyConfigurationList(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/loyalty/configuration/list")
	}
}

func UpdateLoyaltyConfigurationV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodPut:
		helpers.UpdateLoyaltyConfiguration(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/loyalty/configuration/update")
	}
}
