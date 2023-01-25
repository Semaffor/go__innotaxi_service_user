package main

import (
	innotaxi "github.com/Semaffor/go__innotaxi_service_user"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/handler"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository"
	repositoryMongo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongodb"
	repositoryPostgres "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
	serviceMongo "github.com/Semaffor/go__innotaxi_service_user/pkg/service/mongodb"
	servicePostgres "github.com/Semaffor/go__innotaxi_service_user/pkg/service/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Can't read config file: %s", err.Error())
	}
	// postgres, mongo := initDataBaseConnections()
	// handlers := initService(postgres, mongo)
	handlers := handler.NewHandler(nil, nil)

	server := new(innotaxi.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initService(dbPostgre, dbMongo *sqlx.DB) *handler.Handler {
	repoMongo := repositoryMongo.NewRepositoryMongo(dbMongo)
	repoPostgres := repositoryPostgres.NewRepositoryPostgres(dbPostgre)
	serviceMongo := serviceMongo.NewServiceMongo(repoMongo)
	servPostgre := servicePostgres.NewServicePostgre(repoPostgres)

	return handler.NewHandler(serviceMongo, servPostgre)
}

func initDataBaseConnections() (postgre, mongo *sqlx.DB) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Faild to load env data: %s", err.Error())
	}

	newMongo, err := repositoryMongo.NewMongo(repository.Config{
		Host:     viper.GetString("postgres.port"),
		Port:     viper.GetString("postgres.host"),
		Username: viper.GetString("postgres.username"),
		DbName:   viper.GetString("postgres.dbname"),
		SslMode:  viper.GetString("postgres.sslmode"),
		Password: os.Getenv("DB_POSTGRES_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("Can't connect to mongoDB: %s", err.Error())
	}

	newPostgres, err := repositoryMongo.NewMongo(repository.Config{
		Host:     viper.GetString("mongodb.port"),
		Port:     viper.GetString("mongodb.host"),
		Username: viper.GetString("mongodb.username"),
		DbName:   viper.GetString("mongodb.dbname"),
		SslMode:  viper.GetString("mongodb.sslmode"),
		Password: os.Getenv("DB_MONGO_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("Can't connect to mongoDB: %s", err.Error())
	}

	return newPostgres, newMongo
}
