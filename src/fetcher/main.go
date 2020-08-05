package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cryptos, err := makeRequest()
	if err != nil {
		log.Fatal(err)
	}

	client, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://database:27017").SetAuth(
			options.Credential{
				AuthSource: os.Getenv("MONGO_INITDB_DATABASE"),
				Username:   os.Getenv("MONGO_USER"),
				Password:   os.Getenv("MONGO_USER_PWD"),
			}))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	db := client.Database(os.Getenv("MONGO_INITDB_DATABASE"))

	for _, crypto := range cryptos {
		_, err := db.Collection("cryptocurrency").InsertOne(ctx, crypto)
		if err != nil {
			log.Fatal(err)
		}
	}
}
