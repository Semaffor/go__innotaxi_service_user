package log

import (
	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongo"
)

type Logger interface{}

type Service struct {
	Logger
}

func NewService(repo *repo.LogsRepo) *Service {
	return &Service{}
}
