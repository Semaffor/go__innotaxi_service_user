package mongo

import (
	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
)

func NewConnection(config *config.DBConfig) *sqlx.DB {
	// connection to db
	return nil
}
