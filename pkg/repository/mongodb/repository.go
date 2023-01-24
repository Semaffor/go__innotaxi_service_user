package mongodb

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type RepositoryMongo struct {
	Authorization
}

func NewRepositoryMongo(db *sqlx.DB) *RepositoryMongo {
	return &RepositoryMongo{}
}
