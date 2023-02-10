package model

import (
	"github.com/guregu/null"
)

// Roles.
const (
	USER    = 1
	DRIVER  = 2
	ANALYST = 3
)

type User struct {
	Id           int         `db:"id"`
	Name         string      `db:"name"`
	Username     null.String `db:"username"`
	PhoneNumber  string      `db:"phone_number"`
	Email        string      `db:"email"`
	PasswordHash string      `db:"password_hash"`
	Role         int         `db:"role"` // Better int - safe disk capacity or string - increase readability?
	TotalMark    null.Float  `db:"total_mark"`
	IsDeleted    bool        `db:"is_deleted"`
}
