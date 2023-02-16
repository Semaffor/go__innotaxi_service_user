package redis

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/redis/model"
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
	userID int,
	tknID string,
	expiresIn time.Duration,
) error {
	key := fmt.Sprintf("%d:%s", userID, tknID)
	if err := r.db.Set(ctx, key, 0, expiresIn).Err(); err != nil {
		log.Printf("Could not SET refresh token to redis for userID/tknID: %d/%s: %v\n", userID, tknID, err)
		return err
	}

	return nil
}

func (r *TokenRepository) GetByRefreshToken(ctx context.Context, refreshToken string) (*model.Record, error) {
	keyPattern := fmt.Sprintf("*:%s", refreshToken)
	return r.GetByKey(ctx, keyPattern)
}

func (r *TokenRepository) GetByKey(ctx context.Context, keyPattern string) (*model.Record, error) {
	iterator := r.FindKeysByPattern(ctx, keyPattern)

	if !iterator.Next(ctx) {
		return nil, fmt.Errorf("keyPattern %s doesn't exist", keyPattern)
	}
	if err := iterator.Err(); err != nil {
		return nil, err
	}

	key := iterator.Val()
	val, err := r.db.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("keyPattern %s not found", keyPattern)
	}

	splitKey := strings.Split(key, ":")

	return &model.Record{
		UserId: splitKey[0],
		Role:   val,
	}, nil
}

func (r *TokenRepository) DeleteRefreshToken(ctx context.Context, userID int, tokenID string) error {
	key := fmt.Sprintf("%d:%s", userID, tokenID)
	result, err := r.db.Del(ctx, key).Result()
	if result == 0 {
		log.Printf("Trying to delete unexisting key/s: %s", key)
	}

	if err != nil {
		log.Printf("Could not delete refresh token to redis for pattern: %s: %v\n", key, err)
		return err
	}

	return nil
}

func (r *TokenRepository) DeleteAllUserRefreshTokens(ctx context.Context, userID int) error {
	key := fmt.Sprintf("%d:*", userID)
	iterator := r.FindKeysByPattern(ctx, key)

	var foundedRecordCount int = 0
	for iterator.Next(ctx) {
		log.Printf("Deleted= %s\n", iterator.Val())
		r.db.Del(ctx, iterator.Val())
		foundedRecordCount++
	}
	if err := iterator.Err(); err != nil {
		return err
	}
	log.Printf("Deleted Count %d\n", foundedRecordCount)

	return nil
}

func (r *TokenRepository) FindKeysByPattern(ctx context.Context, pattern string) *redis.ScanIterator {
	return r.db.Scan(ctx, 0, pattern, 0).Iterator()
}
