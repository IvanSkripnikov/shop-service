package main

import (
	"context"
	"loyalty_system/helpers"
	"loyalty_system/httphandler"
	"loyalty_system/models"

	logger "github.com/IvanSkripnikov/go-logger"
	migrator "github.com/IvanSkripnikov/go-migrator"
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
	helpers.InitDatabase(config.Database)

	// настройка коннекта к redis
	helpers.InitRedis(context.Background(), config.Redis)

	// выполнение миграций
	migrator.CreateTables(helpers.DB)

	// инициализация REST-api
	httphandler.InitHTTPServer()

	logger.Info("Service started")
}
