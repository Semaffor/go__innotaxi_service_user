package postgres

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
)

type UserService struct {
	repo *postgres.UserRepository
}

func NewUserService(repo *postgres.UserRepository) *UserService {
	return &UserService{repo: repo}
}
