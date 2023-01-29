package mongodb

import (
	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
	repo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository"
)

func NewMongo(connectionCofig *config.ConfigDb) (*sqlx.DB, error) {
	return repo.NewDbConnection(connectionCofig, "mongodb")
}
