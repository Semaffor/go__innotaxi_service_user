package redis

import (
	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/redis"
)

type Authorization interface{}

type ServiceRedis struct {
	Authorization
}

func NewServiceRedis(repo *repo.TokenRepository) *ServiceRedis {
	return &ServiceRedis{}
}
