package helpers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"loyalty_system/models"

	"github.com/IvanSkripnikov/go-logger"

	"github.com/redis/go-redis/v9"
)

func GetItemsList(w http.ResponseWriter, _ *http.Request, user models.User) {
	category := "/v1/items/list"
	var items []models.Item

	err := GormDB.Where("user_category_id <= ?", user.CategoryID).Find(&items).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": items,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	category := "/v1/item/get"
	var item models.Item

	itemID, _ := getIDFromRequestString(strings.TrimSpace(r.URL.Path))

	err := GormDB.Where("id = ?", itemID).First(&item).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": item,
	}
	SendResponse(w, data, "/v1/items/get", http.StatusOK)
}

func CreateItem(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/items/create"
	var item models.Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if checkError(w, err, category) {
		return
	}

	// проверяем, есть ли права на создание товара
	if user.CategoryID == models.UserCategoryStandart {
		FormatResponse(w, http.StatusForbidden, category)
	}

	item.Created = GetCurrentDate()
	item.Updated = GetCurrentDate()

	err = GormDB.Create(&item).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": models.Success,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func UpdateItem(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/items/update"
	var itemRequest models.Item

	err := json.NewDecoder(r.Body).Decode(&itemRequest)
	if checkError(w, err, category) {
		return
	}

	// проверяем, есть ли права на изменение товара
	if user.CategoryID == models.UserCategoryStandart {
		FormatResponse(w, http.StatusForbidden, category)
	}

	var item models.Item
	err = GormDB.Where("id = ?", itemRequest.ID).First(&item).Error
	if checkError(w, err, category) {
		return
	}

	err = GormDB.Model(&item).Updates(models.Item{
		Title:          itemRequest.Title,
		Description:    itemRequest.Description,
		Active:         itemRequest.Active,
		UserCategoryID: itemRequest.UserCategoryID,
		Price:          itemRequest.Price,
		Updated:        GetCurrentDate(),
	}).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": models.Success,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func RemoveItem(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/items/remove"

	// проверяем, есть ли права на удаление товара
	if user.CategoryID == models.UserCategoryStandart {
		FormatResponse(w, http.StatusForbidden, category)
	}

	itemID, err := getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if checkError(w, err, category) {
		return
	}

	err = GormDB.Delete(&models.Item{}, itemID).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": models.Success,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func BuyItem(w http.ResponseWriter, r *http.Request, user models.User) {
	category := "/v1/items/buy"

	RequestID := r.Header.Get("X-Request-Id")
	if RequestID == "" {
		logger.Fatal("Not set X-Request-Id header")
		data := ResponseData{
			"response": models.Failure,
		}
		SendResponse(w, data, "/v1/items/create", http.StatusOK)
		return
	}

	var itemRequest models.BuyItem
	err := json.NewDecoder(r.Body).Decode(&itemRequest)
	if checkError(w, err, category) {
		return
	}

	// 1. Достаём товар из базы
	var item models.Item
	err = GormDB.Where("id = ?", itemRequest.ID).First(&item).Error
	if checkError(w, err, category) {
		logger.Warningf("Not found item: %v", item.ID)
		return
	}

	// 2. Проверяем, доступен ли товар пользователю
	if user.CategoryID < item.UserCategoryID {
		FormatResponse(w, http.StatusForbidden, "/v1/items/buy")
	}

	// 3. Оформляем заказ в сервисе заказов
	response := models.Success
	newOrder := models.Order{UserID: user.ID, ItemID: item.ID, Volume: itemRequest.Volume, Price: item.Price * float32(itemRequest.Volume), RequestID: RequestID}
	newOrderResponse, err := CreateQueryWithResponse(http.MethodPost, Config.OrdersServiceUrl+"/v1/orders/create", newOrder)
	if checkError(w, err, category) || newOrderResponse != models.Success {
		response = models.Failure
		messageData := map[string]interface{}{
			"title":       "Failure buy item!",
			"description": "Something wrong happened during handle your bought: " + item.Title,
			"user":        user.ID,
			"category":    "deal",
		}
		SendNotification(messageData)
	} else {
		messageData := map[string]interface{}{
			"title":       "Successfully buy item!",
			"description": "You successfully buy " + item.Title,
			"user":        user.ID,
			"category":    "deal",
		}
		SendNotification(messageData)
	}

	data := ResponseData{
		"response": response,
	}
	SendResponse(w, data, "/v1/items/create", http.StatusOK)
}

func SendNotification(message map[string]interface{}) {
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		logger.Fatalf("Error connection to Redis: %v", err)
	}

	_, err = redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Stream: Config.Redis.Stream,
		Values: message,
	}).Result()
	if err != nil {
		logger.Fatalf("Error sending to redis stream: %v", err)
	} else {
		logger.Info("Succsessfuly send to stream")
	}
}
