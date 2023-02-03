package mongodb

import (
	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongodb"
)

type Logger interface{}

type Service struct {
	Logger
}

func NewServiceMongo(repo *repo.LogsRepo) *Service {
	return &Service{}
}
