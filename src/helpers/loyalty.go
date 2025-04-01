package helpers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"loyalty_system/models"

	"github.com/IvanSkripnikov/go-logger"
)

func GetLoyaltyList(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/loyalty/list"

	// проверяем, есть ли права на просмотр лояльностей
	if user.CategoryID == models.UserCategoryStandart {
		FormatResponse(w, http.StatusForbidden, category)
		return
	}

	loyalty, err := CreateQueryWithResponse(http.MethodGet, Config.LoyaltyServiceUrl+"/v1/loyalty/list", nil)
	if checkError(w, err, category) {
		logger.Errorf("Error while get loyalty list: %v", err)
	}

	data := ResponseData{
		"response": loyalty,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func GetLoyalty(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/loyalty/get"

	// проверяем, есть ли права на просмотр лояльностей
	if user.CategoryID == models.UserCategoryStandart {
		FormatResponse(w, http.StatusForbidden, category)
		return
	}

	loyaltyID, err := getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if checkError(w, err, category) {
		return
	}

	loyalty, err := CreateQueryWithResponse(http.MethodGet, Config.LoyaltyServiceUrl+"/v1/loyalty/get/"+strconv.Itoa(loyaltyID), nil)
	if checkError(w, err, category) {
		logger.Errorf("Error while get loyalty list: %v", err)
	}

	data := ResponseData{
		"response": loyalty,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func GetLoyaltyForUser(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/loyalty/get-for-user"

	// проверяем, есть ли права на просмотр лояльностей
	if user.CategoryID == models.UserCategoryStandart {
		FormatResponse(w, http.StatusForbidden, category)
		return
	}

	userID, err := getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if checkError(w, err, category) {
		return
	}

	loyalty, err := CreateQueryWithResponse(http.MethodGet, Config.LoyaltyServiceUrl+"/v1/loyalty/get-for-user/"+strconv.Itoa(userID), nil)
	if checkError(w, err, category) {
		logger.Errorf("Error while get loyalty: %v", err)
	}

	data := ResponseData{
		"response": loyalty,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func CreateLoyalty(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/loyalty/create"

	// проверяем, есть ли права на создание лояльности
	if user.CategoryID == models.UserCategoryStandart {
		FormatResponse(w, http.StatusForbidden, category)
		return
	}

	var loyalty models.Loyalty

	err := json.NewDecoder(r.Body).Decode(&loyalty)
	if checkError(w, err, category) {
		return
	}

	loyalty.ManagerID = user.ID
	response, err := CreateQueryWithResponse(http.MethodPost, Config.LoyaltyServiceUrl+"/v1/loyalty/create", loyalty)
	if checkError(w, err, category) {
		logger.Errorf("Error while get create loyalty: %v", err)
	}

	data := ResponseData{
		"response": response,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func UpdateLoyalty(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/loyalty/update"
	var loyalty models.Loyalty

	// проверяем, есть ли права на изменение лояльности
	if user.CategoryID == models.UserCategoryStandart {
		FormatResponse(w, http.StatusForbidden, category)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&loyalty)
	if checkError(w, err, category) {
		return
	}

	response, err := CreateQueryWithResponse(http.MethodPut, Config.LoyaltyServiceUrl+"/v1/loyalty/update", loyalty)
	if checkError(w, err, category) {
		logger.Errorf("Error while get update loyalty: %v", err)
	}

	data := ResponseData{
		"response": response,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func RemoveLoyalty(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/loyalty/remove"

	// проверяем, есть ли права на удаление лояльности
	if user.CategoryID == models.UserCategoryStandart {
		FormatResponse(w, http.StatusForbidden, category)
		return
	}

	loyaltyID, err := getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if checkError(w, err, category) {
		return
	}

	response, err := CreateQueryWithResponse(http.MethodDelete, Config.LoyaltyServiceUrl+"/v1/loyalty/remove/"+strconv.Itoa(loyaltyID), nil)
	if checkError(w, err, category) {
		logger.Errorf("Error while remove loyalty: %v", err)
	}

	data := ResponseData{
		"response": response,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func GetLoyaltyConfigurationList(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/loyalty/configuration/list"

	// проверяем, есть ли права на просмотр лояльностей
	if user.CategoryID == models.UserCategoryStandart {
		FormatResponse(w, http.StatusForbidden, category)
		return
	}

	loyalty, err := CreateQueryWithResponse(http.MethodGet, Config.LoyaltyServiceUrl+"/v1/loyalty/configuration/list", nil)
	if checkError(w, err, category) {
		logger.Errorf("Error while get loyalty list: %v", err)
	}

	data := ResponseData{
		"response": loyalty,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func UpdateLoyaltyConfiguration(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/loyalty/configuration/update"
	var loyaltyConfiguration models.LoyaltyConfiguration

	// проверяем, есть ли права на изменение лояльности
	if user.CategoryID == models.UserCategoryStandart {
		FormatResponse(w, http.StatusForbidden, category)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&loyaltyConfiguration)
	if checkError(w, err, category) {
		return
	}

	response, err := CreateQueryWithResponse(http.MethodPut, Config.LoyaltyServiceUrl+"/v1/loyalty/configuration/update", loyaltyConfiguration)
	if checkError(w, err, category) {
		logger.Errorf("Error while get update loyalty: %v", err)
	}

	data := ResponseData{
		"response": response,
	}
	SendResponse(w, data, category, http.StatusOK)
}
