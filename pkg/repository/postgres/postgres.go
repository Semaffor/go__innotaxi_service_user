package postgres

import (
	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository"
)

func newPostgres(connectionCofig *config.ConfigDb) (*sqlx.DB, error) {
	return repo.NewDbConnection(connectionCofig, "postgres")
}
