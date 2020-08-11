package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoClient() (*mongo.Client, error) {
	return mongo.Connect(context.TODO(),
		options.Client().ApplyURI("mongodb://database:27017").SetAuth(
			options.Credential{
				AuthSource: os.Getenv("MONGO_INITDB_DATABASE"),
				Username:   os.Getenv("MONGO_USER"),
				Password:   os.Getenv("MONGO_USER_PWD"),
			}))
}

// getDatabase will return the database object once the mongo connection is done
func getDatabase() (*mongo.Database, error) {
	client, err := getMongoClient()
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return client.Database(os.Getenv("MONGO_INITDB_DATABASE")), nil
}
