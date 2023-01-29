package redis

import (
	"os"
	"time"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/auth"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/domain"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/helpers"
	repoRedis "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/redis"
)

var (
	ttl = time.Duration(helpers.ConvertToInt(os.Getenv("TOKEN_TTL_MIN"), 20)) * time.Minute
)

type SessionService struct {
	*auth.Manager
	*repoRedis.RepositoryRedis
}

func NewSessionService(repositoryRedis *repoRedis.RepositoryRedis, manager *auth.Manager) *SessionService {
	return &SessionService{Manager: manager, RepositoryRedis: repositoryRedis}
}

func (s *SessionService) CreateSession(user *domain.User) (domain.JwtTokens, error) {
	var (
		tokens domain.JwtTokens
		err    error
	)

	tokens.AccessToken, err = s.Manager.NewJwt(user.Id, user.Username, ttl)
	if err != nil {
		return tokens, err
	}

	tokens.RefreshToken, err = s.Manager.NewRefreshToken()
	if err != nil {
		return tokens, err
	}

	// saveRTinDB

	return tokens, nil
}
