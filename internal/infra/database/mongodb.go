package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBConnection(ctx context.Context, mongoUri string, databaseName string) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	database := client.Database(databaseName)

	return database, nil
}
