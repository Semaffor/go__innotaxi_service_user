package storage

import (
	"context"
)

type Storage interface {
	UserRepo()
	LogsRepo()
	TokenRepo()
}

type UserRepo interface {
	AddUser()
}

type LogsRepo interface {
	Add()
}

type TokenRepo interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}
