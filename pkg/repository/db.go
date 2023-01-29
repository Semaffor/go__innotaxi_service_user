package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
)

func NewDbConnection(c *config.ConfigDb, driverName string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driverName, fmt.Sprintf("host=%s port=%s username=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.Username, c.Password, c.DbName, c.SslMode))

	if err != nil {
		return nil, err
	}

	return db, nil
}
