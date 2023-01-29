package mongodb

import (
	"github.com/jmoiron/sqlx"

	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository"
)

func NewMongo(connectionCofig repo.Config) (*sqlx.DB, error) {
	return repo.NewDbConnection(connectionCofig, "mongodb")
}
