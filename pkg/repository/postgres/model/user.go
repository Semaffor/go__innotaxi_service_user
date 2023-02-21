package model

import (
	"github.com/guregu/null"
)

// Roles.
const (
	USER    = "user"
	DRIVER  = "driver"
	ANALYST = "analyst"
)

type User struct {
	Id           int         `db:"id"`
	Name         string      `db:"name"`
	Username     null.String `db:"username"`
	PhoneNumber  string      `db:"phone_number"`
	Email        string      `db:"email"`
	PasswordHash string      `db:"password_hash"`
	Role         string      `db:"role"`
	TotalMark    null.Float  `db:"total_mark"`
	IsDeleted    bool        `db:"is_deleted"`
}
