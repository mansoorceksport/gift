package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DATABASE = "gift"

func NewMongo(ctx context.Context, connectionString string) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	db := client.Database(DATABASE)
	return db, nil
}
