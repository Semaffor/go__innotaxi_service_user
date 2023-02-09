package model

type Role struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
