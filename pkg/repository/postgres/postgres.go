package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
)

const (
	DriverTable   = "driver"
	FeedbackTable = "feedback"
	OrderTable    = "order"
	RolesTable    = "role"
	TaxiTypeTable = "taxi_type"
	UserRoleTable = "usr_role"
	UsersTable    = "usr"
)

func NewConnection(cfg *config.ConfigDB) *sqlx.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SslMode)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Can't conntect to postgres: %s", err.Error())
	}

	return db
}
