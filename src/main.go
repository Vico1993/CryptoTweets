package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Vico1993/CryptoTweets/src/cmc"
)

func main() {
	cryptoService := cmc.NewCryptoService()

	cryptos, err := cryptoService.GetTopCrypto(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	for _, crypto := range cryptos {
		fmt.Println(crypto.Name)
	}
}
