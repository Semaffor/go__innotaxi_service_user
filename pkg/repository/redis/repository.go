package postgres

import "github.com/jmoiron/sqlx"

type Mock interface {
}

type RepositoryRedis struct {
	Mock
}

func NewRepositoryRedis(db *sqlx.DB) *RepositoryRedis {
	return &RepositoryRedis{}
}
