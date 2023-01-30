package storage

import (
	"context"
)

type Storage interface {
	UserRepo()
	LogsRepo()
	TokenRepo()
}

type UserLayer interface {
	AddUser()
}

type LogsLayer interface {
	Add()
}

type TokenLayer interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}
