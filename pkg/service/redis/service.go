package redis

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/auth"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/domain"
	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/redis"
)

type Authorization interface {
	CreateSession(user *domain.User) (domain.JwtTokens, error)
}

type ServiceRedis struct {
	Authorization
}

func NewServiceRedis(repo *repo.TokenRepo, manager *auth.Manager) *ServiceRedis {
	return &ServiceRedis{
		Authorization: NewSessionService(repo, manager),
	}
}
