package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenRepository struct {
	db *redis.Client
}

func NewTokenRepository(db *redis.Client) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

func (r *TokenRepository) SetRefreshToken(
	ctx context.Context,
	userID string, tknID string,
	expiresIn time.Duration,
) error {
	key := fmt.Sprintf("%s:%s", userID, tknID)
	if err := r.db.Set(ctx, key, 0, expiresIn).Err(); err != nil {
		log.Printf("Could not SET refresh token to redis for userID/tknID: %s/%s: %v\n", userID, tknID, err)

		return err
	}

	return nil
}

func (r *TokenRepository) DeleteRefreshToken(ctx context.Context, userID string, tokenID string) error {
	key := fmt.Sprintf("%s:%s", userID, tokenID)
	result, err := r.db.Del(ctx, key).Result()
	if result == 0 {
		log.Printf("Trying to delete unexisting key: %s", key)
	}

	if err != nil {
		log.Printf("Could not delete refresh token to redis for userID/tokenID: %s/%s: %v\n", userID, tokenID, err)

		return err
	}

	return nil
}
