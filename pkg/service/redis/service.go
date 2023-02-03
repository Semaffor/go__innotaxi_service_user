package redis

import (
	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongodb"
)

type Authorization interface{}

type ServiceRedis struct{}

func NewServiceRedis(repo *repo.LogsRepo) *ServiceRedis {
	return &ServiceRedis{}
}
