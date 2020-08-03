package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type cryptocurrency struct {
	ID                int      `json:"id"`
	Name              string   `json:"name"`
	Symbol            string   `json:"symbol"`
	Slug              string   `json:"slug"`
	CmcRank           float64  `json:"cmc_rank"`
	NumMarketPairs    float64  `json:"num_market_pairs"`
	CirculatingSupply float64  `json:"circulating_supply"`
	TotalSupply       float64  `json:"total_supply"`
	MaxSupply         float64  `json:"max_supply"`
	LastUpdated       string   `json:"last_updated"`
	DateAdded         string   `json:"date_added"`
	Tags              []string `json:"tags"`
}

type cmcOutput struct {
	Data []cryptocurrency `json:"data"`
}

func makeRequest() ([]cryptocurrency, error) {
	config, err := getCMCConfig()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", config.CmcBaseURL+"/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "10")
	q.Add("convert", "CAD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", config.CmcAPI)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	fmt.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		log.Print(resp.Status)
		return nil, errors.New("Request Status code is different that 200, it's : " + string(resp.StatusCode))
	}

	var output cmcOutput
	err = json.Unmarshal(respBody, &output)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return output.Data, nil
}
