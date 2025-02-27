package helpers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"loyalty_system/logger"
	"loyalty_system/models"

	"github.com/redis/go-redis/v9"
)

func GetItemsList(w http.ResponseWriter, _ *http.Request) {
	category := "/v1/items/list"
	var items []models.Item

	query := "SELECT id, title, description, price, created, updated, active FROM items WHERE active = 1"
	rows, err := DB.Query(query)
	if err != nil {
		logger.Error(err.Error())
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err()
	}()

	for rows.Next() {
		item := models.Item{}
		if err = rows.Scan(&item.ID, &item.Title, &item.Description, &item.Price, &item.Created, &item.Updated, &item.Active); err != nil {
			logger.Error(err.Error())
			continue
		}
		items = append(items, item)
	}
	data := ResponseData{
		"data": items,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	category := "/v1/item/get"
	var item models.Item

	item.ID, _ = getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if item.ID == 0 {
		FormatResponse(w, http.StatusUnprocessableEntity, category)
		return
	}

	if !isExists("SELECT * FROM items WHERE id = ?", item.ID) {
		FormatResponse(w, http.StatusNotFound, category)
		return
	}

	query := "SELECT id, title, description, price, created, updated, active FROM items WHERE id = ? AND active = 1"
	rows, err := DB.Prepare(query)

	if checkError(w, err, category) {
		return
	}

	defer func() {
		_ = rows.Close()
	}()

	err = rows.QueryRow(item.ID).Scan(&item.ID, &item.Title, &item.Description, &item.Price, &item.Created, &item.Updated, &item.Active)
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"itemGet": "OK",
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
	// 1. Достаём товар из базы
	var item models.Item

	item.ID, _ = getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if item.ID == 0 {
		FormatResponse(w, http.StatusUnprocessableEntity, category)
		return
	}

	if !isExists("SELECT * FROM items WHERE id = ?", item.ID) {
		FormatResponse(w, http.StatusNotFound, category)
		return
	}

	query := "SELECT id, title, description, price, created, updated, active FROM items WHERE id = ? AND active = 1"
	rows, err := DB.Prepare(query)
	if checkError(w, err, category) {
		return
	}

	defer func() {
		_ = rows.Close()
	}()

	err = rows.QueryRow(item.ID).Scan(&item.ID, &item.Title, &item.Description, &item.Price, &item.Created, &item.Updated, &item.Active)
	if checkError(w, err, category) {
		return
	}

	// 2. Производим списание средств через сервис платежей
	err = WriteOffFromAccount(user.ID, item.Price)
	if checkError(w, err, category) {
		return
	}

	// 3. Оформляем заказ в сервисе заказов
	err = createOrder(user.ID, item.ID, item.Price)
	if checkError(w, err, category) {
		return
	}

	// 4. Информируем клиента об успехе
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		logger.Fatalf("Error connection to Redis: %v", err)
	}

	messageData := map[string]interface{}{
		"title":       "Successfully buy item!",
		"description": "You successfully buy " + item.Title,
		"user":        user.ID,
	}

	_, err = redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Stream: Config.Redis.Stream,
		Values: messageData,
	}).Result()
	if err != nil {
		logger.Fatalf("Error sending to redis stream: %v", err)
	} else {
		logger.Info("Succsessfuly send to stream")
	}

	data := ResponseData{
		"itemCreate": "OK",
	}
	SendResponse(w, data, "/v1/items/create", http.StatusOK)
}

func WriteOffFromAccount(userID int, balance float32) error {
	newAccount := models.Account{UserID: userID, Balance: balance}
	jsonData, err := json.Marshal(newAccount)

	logger.Infof("json for buy: %v", string(jsonData))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, Config.BillingServiceUrl+"/v1/account/buy", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		logger.Fatalf("Error while create PUT deposit request: %v", err)
		return err
	}

	resp, err := client.Do(req)
	logger.Infof("response for make deposit: %v", resp.Body)
	if err != nil {
		logger.Fatalf("Error while process PUT deposit request: %v", err)
		return err
	}

	return nil
}

func DepositForAccount(userID int, balance float32) error {
	newAccount := models.Account{UserID: userID, Balance: balance}
	jsonData, err := json.Marshal(newAccount)
	logger.Infof("json for deposit: %v", string(jsonData))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, Config.BillingServiceUrl+"/v1/account/deposit", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		logger.Fatalf("Error while create PUT deposit request: %v", err)
		return err
	}

	resp, err := client.Do(req)
	logger.Infof("response for make deposit: %v", resp.Body)
	if err != nil {
		logger.Fatalf("Error while process PUT deposit request: %v", err)
		return err
	}

	defer resp.Body.Close()

	return nil
}

func createOrder(userID, itemID int, price float32) error {
	newOrder := models.Order{UserID: userID, ItemID: itemID, Price: price}
	jsonData, err := json.Marshal(newOrder)
	if err != nil {
		return err
	}
	// Отправляем POST-запрос
	resp, err := http.Post(Config.OrdersServiceUrl+"/v1/orders/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result map[string]string

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Infof("Can't parse order data %v", err)
		return err
	}
	json.Unmarshal(body, &result)

	logger.Infof("Data from create payment %v", result)

	// Преобразуем JSON-строку в map
	if err != nil {
		logger.Fatalf("Error convert JSON: %v", err)
		return err
	}

	if result["status"] != "Success" {
		return errors.New("Failed to create order")
	}

	return nil
}
