package mongodb

import (
	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongodb"
)

type Logger interface {
}

type ServiceMongo struct {
	Logger
}

func NewServiceMongo(repo *repo.LogsRepo) *ServiceMongo {
	return &ServiceMongo{}
}
