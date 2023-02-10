package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"

	mongoRepo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongo"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongo/domain"
	postgresRepo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
	redisRepo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/redis"
)

type Repositories struct {
	Users
	Logs
	Tokens
}

func NewRepositories(postgresDB *sqlx.DB, mongoDB *mongo.Database, redisDB *redis.Client) *Repositories {
	return &Repositories{
		Users:  postgresRepo.NewUserRepository(postgresDB),
		Logs:   mongoRepo.NewLogsRepository(mongoDB),
		Tokens: redisRepo.NewTokenRepository(redisDB),
	}
}

type Users interface {
	Save(ctx context.Context, user *model.User) (int, error)
	Update(ctx context.Context, user *model.User) error
	DeleteUserById(ctx context.Context, userId int) error
	FindByFields(ctx context.Context, params map[string]interface{}) ([]model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
}

type Logs interface {
	Write(ctx context.Context, log domain.Log) error
	ReadAllLimit(ctx context.Context, limit int64) ([]domain.Log, error)
	ReadAllByLevel(ctx context.Context, level string) ([]domain.Log, error)
}

type Tokens interface {
	SetRefreshToken(ctx context.Context, userID int, tokenID string, expiresIn time.Duration) error
	DeleteRefreshToken(ctx context.Context, userID int, tokenID string) error
	DeleteAllUserRefreshTokens(ctx context.Context, userID int) error
}
