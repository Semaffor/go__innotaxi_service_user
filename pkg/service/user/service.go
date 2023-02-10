package user

import (
	modelJwt "github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt/model"
	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

type User interface {
	Authentication(credentials *modelJwt.UserCredentials) (model.User, error)
}

type Service struct {
	User
}

func NewService(repo *repo.UserRepository) *Service {
	return &Service{
		User: NewUserService(repo),
	}
}
