package service

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/hash"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository"
)

type Services struct {
	users  User
	token  Tokens
	logger Logger
}

func (s *Services) UserService() User {
	return s.users
}

func (s *Services) LogService() Logger {
	return s.logger
}

func (s *Services) TokenService() Tokens {
	return s.token
}

type Deps struct {
	Repos        *repository.Repositories
	TokenManager jwt.Manager
	Hasher       hash.PasswordHasher
}

func NewServices(deps *Deps) *Services {
	tokenService := NewSessionService(deps.Repos.Tokens, &deps.TokenManager)
	userService := NewUserService(deps.Repos.Users, deps.Hasher)
	logService := NewLogsService(deps.Repos.Logs)

	return &Services{
		users:  userService,
		token:  tokenService,
		logger: logService,
	}
}
