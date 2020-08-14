package cmc

import (
	"context"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

const (
	topCryptoURL = "/cryptocurrency/listings/latest"
)

// CryptoService to get crypto information
type CryptoService struct{}

// Cryptocurrency format sent from cmc
type Cryptocurrency struct {
	ID                int
	Name              string
	Symbol            string
	Slug              string
	CmcRank           float64
	NumMarketPairs    float64
	CirculatingSupply float64
	TotalSupply       float64
	MaxSupply         float64
	LastUpdated       string
	DateAdded         string
	Tags              []string
}

func transform(data []interface{}) []Cryptocurrency {
	cryptos := make([]Cryptocurrency, len(data))
	for i, row := range data {
		var crypto *Cryptocurrency

		mapstructure.Decode(row, &crypto)

		cryptos[i] = *crypto
	}

	return cryptos
}

// NewCryptoService Create a crypto service
func NewCryptoService() *CryptoService {
	return &CryptoService{}
}

// GetTopCrypto return the top of crypto
func (s *CryptoService) GetTopCrypto(ctx context.Context) ([]Cryptocurrency, error) {
	client := NewClient(http.DefaultClient)

	req, err := client.NewRequest(ctx, http.MethodGet, topCryptoURL)
	if err != nil {
		return nil, err
	}

	output, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return transform(output.Data), nil
}
