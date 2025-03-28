package controllers

import (
	"net/http"

	"loyalty_system/helpers"
)

func GetItemsCatogoryListV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		helpers.GetItemsCategoryList(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/item-category/list")
	}
}

func GetItemCategoryV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetItemCategory(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/item-category/get")
	}
}

func CreateItemCategoryV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodPost:
		helpers.CreateItemCategory(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/item-category/create")
	}
}

func UpdateItemCategoryV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodPut:
		helpers.UpdateItemCategory(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/item-category/update")
	}
}

func RemoveItemCategoryV1(w http.ResponseWriter, r *http.Request) {
	auth, user := helpers.GetAuth(r)
	if !auth {
		http.Redirect(w, r, helpers.Config.RedirectUrl+"/signin", http.StatusFound)
		return
	}

	switch r.Method {
	case http.MethodDelete:
		helpers.RemoveItemCategory(w, r, user)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/item-category/remove")
	}
}
