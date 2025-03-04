package main

import (
	"context"
	"loyalty_system/helpers"
	"loyalty_system/httphandler"
	"loyalty_system/models"

	logger "github.com/IvanSkripnikov/go-logger"
)

func main() {
	logger.Debug("Service starting")

	// регистрация общих метрик
	helpers.RegisterCommonMetrics()

	// настройка всех конфигов
	config, err := models.LoadConfig()
	if err != nil {
		logger.Fatalf("Config error: %v", err)
	}

	helpers.InitConfig(config)

	// настройка коннекта к БД
	_, err = helpers.InitDataBase()
	if err != nil {
		logger.Fatalf("Cant initialize DB: %v", err)
	}

	// настройка коннекта к redis
	//bus := events.MakeBus()
	//go helpers.Listen(bus)
	helpers.InitRedis(context.Background(), config.Redis)

	// выполнение миграций
	helpers.CreateTables()

	// инициализация REST-api
	httphandler.InitHTTPServer()

	logger.Info("Service started")
}
