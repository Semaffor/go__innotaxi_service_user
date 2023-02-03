package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
)

func NewConnection(config *config.ConfigDB) *sqlx.DB {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s username=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName, config.SslMode))
	if err != nil {
		log.Fatalf("Can't connect to mongoDB: %s", err.Error())
	}

	return db
}
