package main

import (
	"fmt"
	"loyalty_system/helpers"
	"loyalty_system/httphandler"
	"loyalty_system/logger"
	"loyalty_system/models"
)

func main() {
	logger.Debug("Service starting")

	// регистрация общих метрик
	helpers.RegisterCommonMetrics()

	// настройка всех конфигов
	config, err := models.LoadConfig()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Config error: %v", err))
	}

	helpers.InitConfig(config)

	// настройка коннекта к БД
	_, err = helpers.InitDataBase()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Cant initialize DB: %v", err))
	}

	// выполнение миграций
	helpers.CreateTables()

	// инициализация REST-api
	httphandler.InitHTTPServer()

	logger.Info("Service started")
}
