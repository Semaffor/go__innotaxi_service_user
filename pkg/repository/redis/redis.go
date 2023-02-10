package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
)

func NewConnection(config *config.ConfigDB) *redis.Client {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Can't connect to redis: %s", err.Error())
	}

	return client
}

func shutdown(redisDB *redis.Client) error {
	if err := redisDB.Close(); err != nil {
		return fmt.Errorf("error closing Redis Client: %w", err)
	}

	return nil
}
