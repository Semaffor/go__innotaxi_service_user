package service

import (
	"context"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt"
	modelJwt "github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt/model"
	form "github.com/Semaffor/go__innotaxi_service_user/pkg/handler/model"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

type Service interface {
	UserService() User
	TokenService() Tokens
	LogService() Logger
}

type Tokens interface {
	AuthManager() *jwt.Manager
	CreateSession(ctx context.Context, userId int, role string) (*modelJwt.JwtTokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (*modelJwt.JwtTokens, error)
	LogoutSingle(ctx context.Context, userId int, refreshToken string) error
}

type User interface {
	Authenticate(ctx context.Context, credentials *form.UserLoginInput) (*model.User, error)
	Register(ctx context.Context, user *form.UserRegistrationInput) error
}

type Logger interface {
	WriteLog(ctx context.Context, log string) error // TODO Struct
}
