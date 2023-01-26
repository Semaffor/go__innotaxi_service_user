package auth

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/domain"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type TokenManager interface {
	NewJwt(userId int, username string, ttl time.Duration) (string, error)
	ParseJwt(jwtToken string) (string, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	signature string
}

func NewManager() *Manager {
	signature := os.Getenv("JWT_SIGNATURE")

	return &Manager{signature: signature}
}

func (m *Manager) NewJwt(userId int, username string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, domain.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId:   userId,
		Username: username,
	})

	return token.SignedString([]byte(m.signature))
}

func (m *Manager) ParseJwt(jwtToken string) (string, error) {
	return "", nil
}
