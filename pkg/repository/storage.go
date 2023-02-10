package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"

	mongoRepo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongo"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongo/domain"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
	redisRepo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/redis"
)

type Repositories struct {
	Users
	Role
	Driver
	Order
	TaxiType
	Feedback
	UserRole
	Logs
	Tokens
}

func NewRepositories(postgresDB *sqlx.DB, mongoDB *mongo.Database, redisDB *redis.Client) *Repositories {
	return &Repositories{
		Users:    postgres.NewUserRepository(postgresDB),
		Role:     postgres.NewRoleRepository(postgresDB),
		Driver:   postgres.NewDriverRepository(postgresDB),
		Order:    postgres.NewOrderRepository(postgresDB),
		TaxiType: postgres.NewTaxiTypeRepository(postgresDB),
		Feedback: postgres.NewFeedbackRepository(postgresDB),
		UserRole: postgres.NewUserRoleRepository(postgresDB),
		Logs:     mongoRepo.NewLogsRepository(mongoDB),
		Tokens:   redisRepo.NewTokenRepository(redisDB),
	}
}

type Users interface {
	Save(ctx context.Context, user *model.User) (int, error)
	Update(ctx context.Context, user *model.User) error
	DeleteUserById(ctx context.Context, userId int) error
	FindByFields(ctx context.Context, params map[string]interface{}) ([]model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
}

type Role interface {
	Save(ctx context.Context, user *model.Role) (int, error)
	Update(ctx context.Context, user *model.Role) error
	FindByFields(ctx context.Context, params map[string]interface{}) ([]model.Role, error)
	FindAll(ctx context.Context) ([]model.Role, error)
}

type Driver interface {
	Save(ctx context.Context, user *model.Driver) (int, error)
	Update(ctx context.Context, user *model.Driver) error
	FindByFields(ctx context.Context, params map[string]interface{}) ([]model.Driver, error)
	FindAll(ctx context.Context) ([]model.Driver, error)
}

type Order interface {
	Save(ctx context.Context, user *model.Order) (int, error)
	Update(ctx context.Context, user *model.Order) error
	FindByFields(ctx context.Context, params map[string]interface{}) ([]model.Order, error)
	FindAll(ctx context.Context) ([]model.Order, error)
}

type TaxiType interface {
	Save(ctx context.Context, user *model.TaxiType) (int, error)
	Update(ctx context.Context, user *model.TaxiType) error
	FindByFields(ctx context.Context, params map[string]interface{}) ([]model.TaxiType, error)
	FindAll(ctx context.Context) ([]model.TaxiType, error)
}

type Feedback interface {
	Save(ctx context.Context, user *model.Feedback) (int, error)
	Update(ctx context.Context, user *model.Feedback) error
	FindByFields(ctx context.Context, params map[string]interface{}) ([]model.Feedback, error)
	FindAll(ctx context.Context) ([]model.Feedback, error)
}

type UserRole interface {
	Save(ctx context.Context, user *model.UserRole) (int, error)
	Update(ctx context.Context, user *model.UserRole) error
	FindByFields(ctx context.Context, params map[string]interface{}) ([]model.UserRole, error)
	FindAll(ctx context.Context) ([]model.UserRole, error)
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
