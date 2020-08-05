package database

import (
	"log"

	"github.com/Vico1993/CryptoTweets/src/cmc"
)

const cryptoCollection = "cryptocurrency"

// SaveCryptoCurrency Will save cryptocurrencies in Database
func SaveCryptoCurrency(cryptos []cmc.Cryptocurrency) (bool, error) {
	db, ctx, err := getDatabase()
	if err != nil {
		return false, err
	}

	for _, crypto := range cryptos {
		_, err := db.Collection(cryptoCollection).InsertOne(ctx, crypto)
		if err != nil {
			log.Fatal(err)
		}
	}

	return true, nil
}
