package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cryptos, err := makeRequest()
	if err != nil {
		log.Fatal(err)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{AuthSource: "cryptotweets", Username: "crypto", Password: "secret"}))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("cryptotweets")

	for _, crypto := range cryptos {
		_, err := db.Collection("cryptocurrency").InsertOne(ctx, crypto)
		if err != nil {
			log.Fatal(err)
		}
	}
}
