package postgres

import (
	"fmt"
	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository"
	"github.com/go-redis/redis/v8"
)

func newRedis(config repo.Config) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       0,
	})
	_, err := client.Ping(client.Context()).Result()

	return client, err
}
