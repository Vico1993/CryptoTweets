package cmc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// BasicOuput reponse structure from CMC
type BasicOuput struct {
	Data []interface{} `json:"data"`
}

// Client CMC Config
type Client struct {
	BaseURL    *url.URL
	apiKey     string
	HTTPClient *http.Client
}

// NewClient is creating the CMC client to interact with CMC
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(os.Getenv("CMC_BASE_URL"))

	c := &Client{
		BaseURL:    baseURL,
		apiKey:     os.Getenv("CMC_API"),
		HTTPClient: httpClient,
	}

	return c
}

// NewRequest will build the request to CMC webesite
func (c *Client) NewRequest(ctx context.Context, method string, urlStr string) (*http.Request, error) {
	u := fmt.Sprintf("%s%s", c.BaseURL.String(), urlStr)

	req, err := http.NewRequest(method, u, nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "10")
	q.Add("convert", "CAD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", c.apiKey)
	req.URL.RawQuery = q.Encode()

	return req, nil
}

// CheckResponse check Response status code
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	return errors.New("Request Status code is different that 200, it's : " + string(r.StatusCode))
}

// Do make a request to CMC api
func (c *Client) Do(ctx context.Context, req *http.Request) (BasicOuput, error) {
	var output BasicOuput
	req = req.WithContext(ctx)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Print(err)
		return output, err
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	err = CheckResponse(resp)
	if err != nil {
		return output, err
	}

	err = json.Unmarshal(respBody, &output)

	if err != nil {
		log.Print(err)
		return output, err
	}

	return output, nil
}
