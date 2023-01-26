package redis

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/auth"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/domain"
	repoRedis "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/redis"
	"time"
)

var (
	// ttl := os.Getenv("TOKEN_TTL_HOURS") * time.Hour()
	ttl = time.Now().Add(2).Hour()
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

	tokens.AccessToken, err = s.Manager.NewJwt(user.Id, user.Username, time.Duration(ttl))

	if err != nil {
		return tokens, err
	}
	// createRefreshToken
	// saveRT

	return tokens, nil
}
