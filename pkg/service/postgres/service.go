package postgres

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/domain"
	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
)

type User interface {
	Authentication(credentials *domain.UserCredentials) (domain.User, error)
}

type Service struct {
	User
}

func NewServicePostgre(repo *repo.UserRepository) *Service {
	return &Service{
		User: NewUserService(repo),
	}
}
