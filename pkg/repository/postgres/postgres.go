package postgres

import (
	"github.com/jmoiron/sqlx"
	repo "go__innotaxi_service_user/pkg/repository"
)

func newPostgres(connectionCofig repo.Config) (*sqlx.DB, error) {
	return repo.NewDbConnection(connectionCofig, "postgres")
}
