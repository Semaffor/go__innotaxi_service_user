package postgres

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
)

type UserService struct {
	repo *postgres.RepositoryPostgres
}

func NewUserService(repo *postgres.RepositoryPostgres) *UserService {
	return &UserService{repo: repo}
}
