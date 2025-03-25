package controllers

import (
	"net/http"

	"loyalty_system/helpers"
)

func GetItemsListV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		helpers.GetItemsList(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/items/list")
	}
}

func GetItemV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetItem(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/items/get")
	}
}

func CreateItemV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodPost:
		helpers.CreateItem(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/items/create")
	}
}

func UpdateItemV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodPut:
		helpers.UpdateItem(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/items/update")
	}
}

func RemoveItemV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodDelete:
		helpers.RemoveItem(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/items/remove")
	}
}

func BuyItemV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodPost:
		helpers.BuyItem(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/items/buy")
	}
}
