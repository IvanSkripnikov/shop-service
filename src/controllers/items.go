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
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
