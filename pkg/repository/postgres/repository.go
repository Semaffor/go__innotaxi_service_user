package postgres

import "github.com/jmoiron/sqlx"

type Authentication interface{}

type Order interface{}

type User interface{}

type UserRepository struct {
	Authentication
	Order
	User
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{}
}
