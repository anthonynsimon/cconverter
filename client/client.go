package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anthonynsimon/cconverter/currency"
)

var (
	APIRequestTimeout = time.Second * 10
)

const (
	RatesUri      = "/api/rates/%s"
	ConversionUri = "/api/convert?from=%s&to=%s&amount=%f"
)

type APIClient struct {
	httpClient *http.Client
	apiHost    string
}

func NewClient(apiHost string) *APIClient {
	var client = http.Client{
		Timeout: APIRequestTimeout,
	}

	return &APIClient{
		httpClient: &client,
		apiHost:    apiHost,
	}
}

func (client *APIClient) getRatesURL(currency currency.Currency) string {
	return fmt.Sprintf(client.apiHost+RatesUri, currency)
}

func (client *APIClient) getConversionURL(from, to currency.Currency, value float64) string {
	return fmt.Sprintf(client.apiHost+ConversionUri, from, to, value)
}
