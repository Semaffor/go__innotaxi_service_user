package app

import (
	"log"

	innotaxi "github.com/Semaffor/go__innotaxi_service_user"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/handler"
	repositoryMongo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongo"
	repositoryPostgres "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
)

func Run() error {
	configs, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Can't read config/env file: %s", err.Error())
	}

	postgres, err := repositoryPostgres.NewConnection(&configs.Postgres)
	if err != nil {
		return err
	}

	mongo, err := repositoryMongo.NewConnection(&configs.Mongo)
	if err != nil {
		return err
	}

	services := initServices(postgres, mongo)
	newHandler := handler.NewHandler(services)

	server := new(innotaxi.Server)
	if err := server.Run(&configs.Server, newHandler.InitRoutes()); err != nil {
		log.Println("Error occurred while running.")
		return err
	}

	return nil
}
