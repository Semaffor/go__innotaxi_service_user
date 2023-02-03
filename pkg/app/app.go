package app

import (
	"log"

	"github.com/joho/godotenv"

	innotaxi "github.com/Semaffor/go__innotaxi_service_user"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/handler"
	repositoryMongo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongodb"
	repositoryPostgres "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
)

const configDir = "configs"

func Run() error {
	if err := initConfig(configDir); err != nil {
		log.Fatalf("Can't read configurator file.")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Faild to load env data: %s", err.Error())
	}

	configPostgres := config.ReadConfig("postgres", &config.ConfigDB{})
	postgres := repositoryPostgres.NewConnection(configPostgres)

	configMongo := config.ReadConfig("mongo", &config.ConfigDB{})
	mongo := repositoryMongo.NewConnection(configMongo)

	services := initServices(postgres, mongo)
	newHandler := handler.NewHandler(services)

	server := new(innotaxi.Server)
	serverConfig := config.ReadConfig("server", &config.ServerConfig{})
	if err := server.Run(serverConfig, newHandler.InitRoutes()); err != nil {
		log.Println("Error occurred while running.")
		return err
	}

	return nil
}
