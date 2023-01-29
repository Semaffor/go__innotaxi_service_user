package postgres

import (
	"github.com/jmoiron/sqlx"

	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository"
)

func newPostgres(connectionCofig repo.Config) (*sqlx.DB, error) {
	return repo.NewDbConnection(connectionCofig, "postgres")
}
