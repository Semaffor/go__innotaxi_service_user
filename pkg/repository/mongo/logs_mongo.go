package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/mongo/domain"
)

const (
	logsCollection = "logs"
)

type LogsRepo struct {
	db *mongo.Collection
}

func NewLogsRepository(mongo *mongo.Database) *LogsRepo {
	return &LogsRepo{db: mongo.Collection(logsCollection)}
}

func (r *LogsRepo) Write(ctx context.Context, log domain.Log) error {
	_, err := r.db.InsertOne(ctx, log)
	if err != nil {
		return err
	}

	return nil
}

func (r *LogsRepo) ReadAllLimit(ctx context.Context, limit int64) ([]domain.Log, error) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"createdOn", -1}}).SetLimit(limit)

	return r.findByFilterOrOpts(ctx, filter, opts)
}

func (r *LogsRepo) ReadAllByLevel(ctx context.Context, level string) ([]domain.Log, error) {
	filter := bson.D{bson.E{Key: "level", Value: level}}

	return r.findByFilterOrOpts(ctx, filter, nil)
}

func (r *LogsRepo) findByFilterOrOpts(
	ctx context.Context,
	filter bson.D,
	opts *options.FindOptions,
) ([]domain.Log, error) {
	var content []domain.Log
	cur, err := r.db.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &content); err != nil {
		return nil, err
	}

	return content, nil
}
