package model

import (
	"fmt"

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
	Role         string      `db:"role"`
	TotalMark    null.Float  `db:"total_mark,omitempty"`
	IsDeleted    bool        `db:"is_deleted"`
}

// Custom output - probably, temporary.
func (u *User) String() string {
	return fmt.Sprintf("%d %s %s", u.Id, u.Name, u.Username.String)
}
