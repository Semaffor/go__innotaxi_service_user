package token

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt"
	modelJwt "github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt/model"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/redis"
)

type Authorization interface {
	GetAuthManager() *jwt.Manager
	CreateSession(user *model.User) (modelJwt.JwtTokens, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repo.TokenRepository, manager *jwt.Manager) *Service {
	return &Service{
		Authorization: NewSessionService(repo, manager),
	}
}
