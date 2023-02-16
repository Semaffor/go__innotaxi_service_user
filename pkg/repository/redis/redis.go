package redis

import (
	"fmt"

	rediska "github.com/go-redis/redis/v8"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
)

func NewConnection(config *config.DBConfig) (*rediska.Client, error) {
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	client := rediska.NewClient(&rediska.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       0,
	})
	_, err := client.Ping(client.Context()).Result()

	return client, err
}
