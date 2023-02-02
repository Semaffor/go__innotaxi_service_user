package mongodb

import (
	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongodb"
)

type Authorization interface {
}

type ServiceMongo struct {
	//
	Authorization
}

func NewServiceRedis(repo *repo.LogsRepo) *ServiceMongo {
	return &ServiceMongo{}
}
