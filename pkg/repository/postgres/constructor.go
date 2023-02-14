package postgres

import (
	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/general"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

const (
	UsersTable = "usr"
)

func NewUserDao(db *sqlx.DB) general.DaoI[model.User] {
	return &general.Dao[model.User]{
		Db:    db,
		Table: UsersTable,
	}
}
