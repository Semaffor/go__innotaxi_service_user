package postgres

import (
	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
)

type User interface {
}

type ServicePostgres struct {
	User
}

func NewServicePostgre(repo *repo.UserRepository) *ServicePostgres {
	return &ServicePostgres{
		User: NewUserService(repo),
	}
}
