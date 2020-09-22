package fetcher

import (
	"context"
	"fmt"
	"log"

	"github.com/Vico1993/CryptoTweets/src/cmc"
)

// test was made to implement the fetcher. Move from main to here to not loose this code right now. but it's not needed at the moment
func test() {
	cryptoService := cmc.NewCryptoService()

	cryptos, err := cryptoService.GetTopCrypto(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	for _, crypto := range cryptos {
		fmt.Println(crypto)
	}
}
