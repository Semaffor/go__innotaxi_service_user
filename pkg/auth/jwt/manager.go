package jwt

import (
	"encoding/base32"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt/model"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
)

type TokenManager interface {
	NewJwt(userId int, username string, ttl time.Duration) (string, error)
	ParseJwt(jwtToken string) (string, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	Config    *config.JWTConfig
	Signature string
}

func NewManager(config *config.JWTConfig) *Manager {
	return &Manager{Config: config, Signature: config.SigningKey}
}

func (m *Manager) NewJwt(userId int, role string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
		Role:   role,
	})

	return token.SignedString([]byte(m.Signature))
}

func (m *Manager) ParseJwt(jwtToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (decryptedToken interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.Signature), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("error get user claims from token")
	}

	return map[string]interface{}{
		"userId": claims["userId"],
		"role":   claims["role"],
	}, nil
}

func (m *Manager) NewRefreshToken(tokenLength int) (string, error) {
	randomBytes := make([]byte, 32)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	token := base32.StdEncoding.EncodeToString(randomBytes)[:tokenLength]

	return token, nil
}
