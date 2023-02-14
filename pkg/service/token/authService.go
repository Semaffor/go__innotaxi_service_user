package token

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt"
	modelJwt "github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt/model"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
	repoRedis "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/redis"
)

type SessionService struct {
	*jwt.Manager
	*repoRedis.TokenRepository
}

func NewSessionService(repositoryRedis *repoRedis.TokenRepository, manager *jwt.Manager) *SessionService {
	return &SessionService{Manager: manager, TokenRepository: repositoryRedis}
}

func (s *SessionService) GetAuthManager() *jwt.Manager {
	return s.Manager
}

func (s *SessionService) CreateSession(user *model.User) (modelJwt.JwtTokens, error) {
	var (
		tokens modelJwt.JwtTokens
		err    error
	)

	tokens.AccessToken, err = s.Manager.NewJwt(user.Id, user.Username, s.Config.AccessTokenTTL)
	if err != nil {
		return tokens, err
	}

	tokens.RefreshToken, err = s.Manager.NewRefreshToken(s.Config.RefreshTokenLength)
	if err != nil {
		return tokens, err
	}

	return tokens, nil
}
