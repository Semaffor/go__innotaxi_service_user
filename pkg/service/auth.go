package service

import (
	"context"
	"strconv"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt"
	modelJwt "github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt/model"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository"
)

type SessionService struct {
	manager *jwt.Manager
	repo    repository.Tokens
}

func NewSessionService(repo repository.Tokens, manager *jwt.Manager) *SessionService {
	return &SessionService{manager: manager, repo: repo}
}

func (s *SessionService) AuthManager() *jwt.Manager {
	return s.manager
}

func (s *SessionService) CreateSession(ctx context.Context, userId int, role string) (*modelJwt.JwtTokens, error) {
	var (
		tokens modelJwt.JwtTokens
		err    error
	)

	tokens.AccessToken, err = s.manager.NewJwt(userId, role, s.AuthManager().Config.AccessTokenTTL)
	if err != nil {
		return nil, err
	}

	tokens.RefreshToken, err = s.manager.NewRefreshToken(s.AuthManager().Config.RefreshTokenLength)
	if err != nil {
		return nil, err
	}

	return &tokens, nil
}

func (s *SessionService) RefreshTokens(ctx context.Context, refreshToken string) (*modelJwt.JwtTokens, error) {
	user, err := s.repo.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	userId, err := strconv.Atoi(user.UserId)
	if err != nil {
		return nil, err
	}

	return s.CreateSession(ctx, userId, user.UserId)
}

func (s *SessionService) LogoutSingle(ctx context.Context, userId int, refreshToken string) error {
	return s.repo.DeleteRefreshToken(ctx, userId, refreshToken)
}
