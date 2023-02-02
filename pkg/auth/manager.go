package auth

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/domain"
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

func (m *Manager) NewJwt(userId int, role string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, domain.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
		Role:   role,
	})

	return token.SignedString([]byte(m.signature))
}

func (m *Manager) ParseJwt(jwtToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signature), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("error get user claims from token")
	}

	return map[string]interface{}{
		"userId": claims["userId"].(int),
		"role":   claims["role"].(string),
	}, nil
}

func (m *Manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
