package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
)

func NewConnection(config *config.ConfigDB) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       0,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return client, nil
}

func shutdown(redisDB *redis.Client) error {
	if err := redisDB.Close(); err != nil {
		return fmt.Errorf("error closing Redis Client: %w", err)
	}

	return nil
}
