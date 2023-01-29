package postgres

import "github.com/jmoiron/sqlx"

type Mock interface {
}

type RepositoryPostgres struct {
}

func NewRepositoryPostgres(db *sqlx.DB) *RepositoryPostgres {
	return &RepositoryPostgres{}
}
