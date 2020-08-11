package main

import (
	"log"

	"github.com/Vico1993/CryptoTweets/src/cmc"
	"github.com/Vico1993/CryptoTweets/src/database"
)

func main() {
	cryptos, err := cmc.MakeRequest()
	if err != nil {
		log.Fatal(err)
	}

	database.SaveCryptoCurrency(cryptos)
}
