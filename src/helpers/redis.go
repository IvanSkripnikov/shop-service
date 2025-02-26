package helpers

import (
	"context"
	"net"
	"strconv"

	"loyalty_system/logger"
	"loyalty_system/models"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	cont   context.Context
	stream string
)

// Init Инициализация подключения к Redis.
func InitRedis(ctx context.Context, config models.Redis) {
	if _, err := strconv.Atoi(config.Port); err != nil {
		logger.Fatalf("Failed to parse on Redis port. Error: %v", err)
	}

	address := net.JoinHostPort(config.Address, config.Port)
	client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: config.Password,
		DB:       config.DB,
	})
	cont = ctx
	stream = config.Stream
	logger.Info("Redis initialized")
}

/*
// Listen Прослушивать сообщения в каналах.
func Listen(bus events.EventBus) {
	for {
		select {
		case err := <-bus.Error:
			logger.Error(err.Error())
		}
	}
}

// ListenStream Прослушивание стрима Redis.
func ListenStream(handler func(redis.XMessage), errCh chan<- error) {
	logger.Info("Listening stream...")
	lastId := "0"
	for {
		result, err := client.XRead(cont, &redis.XReadArgs{
			Count:   100,
			Block:   0,
			Streams: []string{stream, lastId},
		}).Result()

		if err != nil {
			logger.Errorf("Cant execute XRead command. Error: %v", err)
			errCh <- err
			return
		}

		messages := result[0].Messages
		countMessages := len(messages)

		if countMessages > 0 {
			logger.Debugf("XRead iteration from ID: %s. New messages: %d", lastId, countMessages)
			lastId = messages[countMessages-1].ID
		}

		for _, message := range messages {
			handler(message)
		}
	}
}
*/
