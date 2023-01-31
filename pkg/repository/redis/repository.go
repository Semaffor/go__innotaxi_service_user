package redis

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type TokenRepo interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type TokenRepository struct {
	TokenRepo
}

func NewTokenRepository(db *sqlx.DB) *TokenRepository {
	return &TokenRepository{}
}
