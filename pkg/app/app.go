package app

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	innotaxi "github.com/Semaffor/go__innotaxi_service_user"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/handler"
	repositoryMongo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongodb"
	repositoryPostgres "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
	serviceMongo "github.com/Semaffor/go__innotaxi_service_user/pkg/service/mongodb"
	servicePostgres "github.com/Semaffor/go__innotaxi_service_user/pkg/service/postgres"
)

const configDir = "configs"

func Run() error {
	if err := initConfig(configDir); err != nil {
		log.Fatalf("Can't read configurator file.")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Faild to load env data: %s", err.Error())
	}

	// init repos
	configPostgres := config.ReadConfig("postgres", &config.ConfigDb{})
	_ = repositoryPostgres.NewConnection(configPostgres)

	configMongo := config.ReadConfig("mongo", &config.ConfigDb{})
	_ = repositoryMongo.NewConnection(configMongo)

	// init services
	//

	handlers := handler.NewHandler(nil, nil)

	server := new(innotaxi.Server)
	serverConfig := config.ReadConfig("server", &config.ServerConfig{})
	if err := server.Run(serverConfig, handlers.InitRoutes()); err != nil {
		log.Println("Error occurred while running.")
		return err
	}

	return nil
}

func initService(dbPostgre, dbMongo *sqlx.DB) *handler.Handler {
	repoMongo := repositoryMongo.NewLogsRepository(dbMongo)
	repoPostgres := repositoryPostgres.NewUserRepository(dbPostgre)
	servMongo := serviceMongo.NewServiceMongo(repoMongo)
	servPostgre := servicePostgres.NewServicePostgre(repoPostgres)

	return handler.NewHandler(servMongo, servPostgre)
}

func initConfig(configDir string) error {
	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
