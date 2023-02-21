package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"

	mongoRepo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongo"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongo/domain"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
	pgModel "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
	redisRepo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/redis"
	redisModel "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/redis/model"
)

type Repositories struct {
	Users
	Logs
	Tokens
}

func NewRepositories(postgresDB *sqlx.DB, mongoDB *mongo.Database, redisDB *redis.Client) *Repositories {
	return &Repositories{
		Users:  postgres.NewUserRepository(postgresDB),
		Logs:   mongoRepo.NewLogsRepository(mongoDB),
		Tokens: redisRepo.NewTokenRepository(redisDB),
	}
}

type Users interface {
	DeleteUserById(ctx context.Context, userId int) error
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*pgModel.User, error)
	FindByUsername(ctx context.Context, phoneNumber string) (*pgModel.User, error)
	FindAll(ctx context.Context) ([]pgModel.User, error)
	Save(ctx context.Context, user *pgModel.User) (int, error)
	Update(ctx context.Context, user *pgModel.User) error
}

type Logs interface {
	ReadAllByLevel(ctx context.Context, level string) ([]domain.Log, error)
	ReadAllLimit(ctx context.Context, limit int64) ([]domain.Log, error)
	Write(ctx context.Context, log domain.Log) error
}

type Tokens interface {
	DeleteAllUserRefreshTokens(ctx context.Context, userID int) error
	DeleteRefreshToken(ctx context.Context, userID int, tokenID string) error
	GetByKey(ctx context.Context, keyPattern string) (*redisModel.Record, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (*redisModel.Record, error)
	SetRefreshToken(ctx context.Context, userID int, tokenID string, expiresIn time.Duration) error
}
