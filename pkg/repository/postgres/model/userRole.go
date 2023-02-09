package model

type UserRole struct {
	id     int `db:"id"`
	userId int `db:"user_id"`
	roleId int `db:"role_id"`
}
