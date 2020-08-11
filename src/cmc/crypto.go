package cmc

import (
	"context"
	"net/http"
	"path"
	"reflect"
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
	for i := range data {
		var crypto Cryptocurrency
		item := reflect.ValueOf(data)

		crypto.ID = item.Field(0).Interface().(int)
		crypto.Name = item.Field(1).Interface().(string)
		crypto.Symbol = item.Field(2).Interface().(string)
		crypto.Slug = item.Field(3).Interface().(string)
		crypto.CmcRank = item.Field(4).Interface().(float64)
		crypto.NumMarketPairs = item.Field(5).Interface().(float64)
		crypto.CirculatingSupply = item.Field(6).Interface().(float64)
		crypto.TotalSupply = item.Field(7).Interface().(float64)
		crypto.MaxSupply = item.Field(8).Interface().(float64)
		crypto.LastUpdated = item.Field(9).Interface().(string)
		crypto.DateAdded = item.Field(10).Interface().(string)
		crypto.Tags = item.Field(11).Interface().([]string)

		cryptos[i] = crypto
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

	req, err := client.NewRequest(ctx, http.MethodGet, path.Join(client.BaseURL.String(), topCryptoURL))
	if err != nil {
		return nil, err
	}

	output, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return transform(output.Data), nil
}
