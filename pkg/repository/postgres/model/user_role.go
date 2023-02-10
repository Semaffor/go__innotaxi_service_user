package model

type UserRole struct {
	Id     int `db:"id"`
	UserId int `db:"user_id"`
	RoleId int `db:"role_id"`
}
