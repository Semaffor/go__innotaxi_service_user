package mongo

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
)

func NewConnection(config *config.ConfigDB) *mongo.Database {
	addr := net.JoinHostPort(config.Host, config.Port)
	uri := fmt.Sprintf("mongodb://%s:%s@%s/", config.Username, config.Password, addr)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Can't connect to mongoDB: %s", err.Error())
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatalf("Can't connect to mongoDB: %s", err.Error())
	}
	db := client.Database(config.DBName)

	return db
}

func shutdown(client *mongo.Client) {
	if err := client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
