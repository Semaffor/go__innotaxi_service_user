package config

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	repositoryMongo "github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongodb"
)

type Server struct {
	Host           string        `mapstructure:"host" default:"localhost" json:"host,omitempty"` // TODO: default not working
	Port           string        `mapstructure:"port" default:"8000" json:"port,omitempty"`
	MaxHeaderBytes int           `mapstructure:"maxHeaderBytes" default:"20"` // number of bit shifts to the left
	ReadTimeout    time.Duration `mapstructure:"readTimeout" default:"10"`    // time in seconds
	WriteTimeout   time.Duration `mapstructure:"writeTimeout" default:"10"`   // time in seconds
}

type ConfigDb struct {
	Host     string `mapstructure:"host" default:"localhost"`
	Port     string `mapstructure:"port" default:"8001"`
	Username string `mapstructure:"username"`
	DbName   string `mapstructure:"dbName"`
	SslMode  string `mapstructure:"sslMode"`
	Password string `mapstructure:"password"`
}

func InitDataBaseConnections() (postgresDb, mongoDb *sqlx.DB) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Faild to load env data: %s", err.Error())
	}

	configPostgres := ReadConfig("postgres", &ConfigDb{})
	configPostgres.Password = os.Getenv("DB_POSTGRES_PASSWORD")
	postgresDb, err := repositoryMongo.NewMongo(configPostgres)
	if err != nil {
		log.Fatalf("Can't connect to mongoDB: %s", err.Error())
	}

	configMongo := ReadConfig("mongo", &ConfigDb{})
	configMongo.Password = os.Getenv("DB_MONGO_PASSWORD")
	mongoDb, err = repositoryMongo.NewMongo(configMongo)
	if err != nil {
		log.Fatalf("Can't connect to mongoDB: %s", err.Error())
	}

	return postgresDb, mongoDb
}
