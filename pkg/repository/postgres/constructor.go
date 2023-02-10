package postgres

import (
	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/general"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

func NewUserDao(db *sqlx.DB) general.DaoI[model.User] {
	return &general.Dao[model.User]{
		Db:    db,
		Table: UsersTable,
	}
}

func NewRoleDao(db *sqlx.DB) general.DaoI[model.Role] {
	return &general.Dao[model.Role]{
		Db:    db,
		Table: RolesTable,
	}
}

func NewDriverDao(db *sqlx.DB) general.DaoI[model.Driver] {
	return &general.Dao[model.Driver]{
		Db:    db,
		Table: DriverTable,
	}
}

func NewOrderDao(db *sqlx.DB) general.DaoI[model.Order] {
	return &general.Dao[model.Order]{
		Db:    db,
		Table: OrderTable,
	}
}

func NewTaxiTypeDao(db *sqlx.DB) general.DaoI[model.TaxiType] {
	return &general.Dao[model.TaxiType]{
		Db:    db,
		Table: TaxiTypeTable,
	}
}

func NewFeedbackDao(db *sqlx.DB) general.DaoI[model.Feedback] {
	return &general.Dao[model.Feedback]{
		Db:    db,
		Table: FeedbackTable,
	}
}

func NewUserRoleDao(db *sqlx.DB) general.DaoI[model.UserRole] {
	return &general.Dao[model.UserRole]{
		Db:    db,
		Table: UserRoleTable,
	}
}
