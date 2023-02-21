package app

import (
	"log"

	innotaxi "github.com/Semaffor/go__innotaxi_service_user"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/handler"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/hash"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository"
	repoMongo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongo"
	repoPostgres "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres"
	repoRedis "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/redis"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service"
)

func Run() error {
	configs, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Can't read config/env file: %s", err.Error())
	}

	repositories, err := initRepositories(configs)
	if err != nil {
		log.Fatalf("Can't init repo layer: %v", err.Error())
	}

	services := initServices(repositories, configs)
	newHandler := handler.NewHandler(services)

	server := new(innotaxi.Server)
	if err := server.Run(&configs.Server, newHandler.InitRoutes()); err != nil {
		log.Println("Error occurred while running.")
		return err
	}

	return nil
}

func initRepositories(configs *config.Config) (*repository.Repositories, error) {
	postgresCon, err := repoPostgres.NewConnection(&configs.Postgres)
	if err != nil {
		return nil, err
	}

	mongoCon, err := repoMongo.NewConnection(&configs.Mongo)
	if err != nil {
		return nil, err
	}

	redisCon, err := repoRedis.NewConnection(&configs.Redis)
	if err != nil {
		return nil, err
	}

	return repository.NewRepositories(postgresCon, mongoCon, redisCon), nil
}

func initServices(repos *repository.Repositories, configs *config.Config) *service.Services {
	deps := &service.Deps{
		Repos:        repos,
		TokenManager: *jwt.NewManager(&configs.AuthConfig.JWT),
		Hasher:       hash.NewSHA256Hasher(configs.AuthConfig.PasswordSalt),
	}

	return service.NewServices(deps)
}
