package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoClient() (*mongo.Client, error) {
	return mongo.NewClient(
		options.Client().ApplyURI("mongodb://database:27017").SetAuth(
			options.Credential{
				AuthSource: os.Getenv("MONGO_INITDB_DATABASE"),
				Username:   os.Getenv("MONGO_USER"),
				Password:   os.Getenv("MONGO_USER_PWD"),
			}))
}

// getDatabase will return the database object once the mongo connection is done
func getDatabase() (*mongo.Database, context.Context, error) {
	client, err := getMongoClient()
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer client.Disconnect(ctx)

	return client.Database(os.Getenv("MONGO_INITDB_DATABASE")), ctx, nil
}
