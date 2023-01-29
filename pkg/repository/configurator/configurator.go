package configurator

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
	repositoryMongo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongodb"
)

func InitDataBaseConnections() (postgresDb, mongoDb *sqlx.DB) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Faild to load env data: %s", err.Error())
	}

	configPostgres := config.ReadConfig("postgres", &config.ConfigDb{})
	configPostgres.Password = os.Getenv("DB_POSTGRES_PASSWORD")
	postgresDb, err := repositoryMongo.NewMongo(configPostgres)
	if err != nil {
		log.Fatalf("Can't connect to mongoDB: %s", err.Error())
	}

	configMongo := config.ReadConfig("mongo", &config.ConfigDb{})
	configMongo.Password = os.Getenv("DB_MONGO_PASSWORD")
	mongoDb, err = repositoryMongo.NewMongo(configMongo)
	if err != nil {
		log.Fatalf("Can't connect to mongoDB: %s", err.Error())
	}

	return postgresDb, mongoDb
}
