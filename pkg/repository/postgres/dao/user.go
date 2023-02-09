package dao

import (
	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/general"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

func NewUserDao(db *sqlx.DB, table string) general.DaoI[model.User] {
	return &general.Dao[model.User]{
		Db:    db,
		Table: table,
	}
}
