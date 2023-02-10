package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"

	repositoryMongo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongo"
	repositoryPostgres "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service"
	serviceMongo "github.com/Semaffor/go__innotaxi_service_user/pkg/service/log"
	servicePostgres "github.com/Semaffor/go__innotaxi_service_user/pkg/service/user"
)

func initServices(postgresDB *sqlx.DB, mongoDB *mongo.Database) *service.Aggregator {
	postgres := initPostgres(postgresDB)
	mongo := initMongo(mongoDB)

	return service.NewAggregator(mongo, postgres, nil)
}

func initConfig(configDir string) error {
	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func initMongo(dbConnection *mongo.Database) *serviceMongo.Service {
	repoMongo := repositoryMongo.NewLogsRepository(dbConnection)

	return serviceMongo.NewService(repoMongo)
}

func initPostgres(dbConnection *sqlx.DB) *servicePostgres.Service {
	repoPostgres := repositoryPostgres.NewUserRepository(dbConnection)

	return servicePostgres.NewService(repoPostgres)
}
