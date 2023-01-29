package redis

import "github.com/jmoiron/sqlx"

type Mock interface {
}

type RepositoryRedis struct {
}

func NewRepositoryRedis(db *sqlx.DB) *RepositoryRedis {
	return &RepositoryRedis{}
}
