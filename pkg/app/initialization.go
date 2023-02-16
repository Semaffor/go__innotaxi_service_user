package app

import (
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"

	repositoryMongo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongo"
	repositoryPostgres "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service"
	serviceMongo "github.com/Semaffor/go__innotaxi_service_user/pkg/service/log"
	servicePostgres "github.com/Semaffor/go__innotaxi_service_user/pkg/service/user"
)

func initServices(dbPostgre *sqlx.DB, dbMongo *mongo.Database) *service.Aggregator {
	postgresCon := initPostgres(dbPostgre)
	mongoCon := initMongo(dbMongo)

	return service.NewAggregator(mongoCon, postgresCon, nil)
}

func initMongo(dbConnection *mongo.Database) *serviceMongo.Service {
	repoMongo := repositoryMongo.NewLogsRepository(dbConnection)

	return serviceMongo.NewService(repoMongo)
}

func initPostgres(dbConnection *sqlx.DB) *servicePostgres.Service {
	repoPostgres := repositoryPostgres.NewUserRepository(dbConnection)

	return servicePostgres.NewService(repoPostgres)
}
