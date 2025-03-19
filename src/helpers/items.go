package helpers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"loyalty_system/models"

	"github.com/IvanSkripnikov/go-gormdb"
	"github.com/IvanSkripnikov/go-logger"

	"github.com/redis/go-redis/v9"
)

func GetItemsList(w http.ResponseWriter, _ *http.Request) {
	category := "/v1/items/list"
	var items []models.Item

	db := gormdb.GetClient(models.ServiceDatabase)
	err := db.Find(&items).Error
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

	db := gormdb.GetClient(models.ServiceDatabase)
	err := db.Where("id = ?", itemID).First(&item).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"itemGet": item,
	}
	SendResponse(w, data, "/v1/items/get", http.StatusOK)
}

func CreateItem(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"itemCreate": "OK",
	}
	SendResponse(w, data, "/v1/items/create", http.StatusOK)
}

func UpdateItem(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"itemUpdate": "OK",
	}
	SendResponse(w, data, "/v1/items/update", http.StatusOK)
}

func RemoveItem(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"itemRemove": "OK",
	}
	SendResponse(w, data, "/v1/items/remove", http.StatusOK)
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
	db := gormdb.GetClient(models.ServiceDatabase)
	err = db.Where("id = ?", itemRequest.ID).First(&item).Error
	if checkError(w, err, category) {
		logger.Warningf("Not found item: %v", item.ID)
		return
	}

	// 2. Оформляем заказ в сервисе заказов
	response := models.Success
	newOrder := models.Order{UserID: user.ID, ItemID: item.ID, Volume: itemRequest.Volume, Price: item.Price * float32(itemRequest.Volume), RequestID: RequestID}
	newOrderResponse, err := CreateQueryWithScalarResponse(http.MethodPost, Config.OrdersServiceUrl+"/v1/orders/create", newOrder)
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
