package mongo

import (
	"context"
	"fmt"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
)

func NewConnection(config *config.ConfigDB) (*mongo.Database, error) {
	addr := net.JoinHostPort(config.Host, config.Port)
	uri := fmt.Sprintf("mongodb://%s:%s@%s/", config.Username, config.Password, addr)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, err
	}
	db := client.Database(config.DBName)

	return db, nil
}

func shutdown(client *mongo.Client) {
	if err := client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
