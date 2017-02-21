package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anthonynsimon/cconverter/currency"
	"github.com/shopspring/decimal"
)

var (
	// APIRequestTimeout is the max amount of time until an API request is cancelled
	APIRequestTimeout = time.Second * 10
)

const (
	// RatesUri is the URI template to the rates endpoint
	RatesUri = "/api/rates/%s"
	// ConversionUri is the URI template to the convert endpoint
	ConversionUri = "/api/convert?from=%s&to=%s&amount=%s"
)

// APIClient holds all methods for interacting with the cconverter API server
type APIClient struct {
	httpClient *http.Client
	apiHost    string
}

// NewClient returns an APIClient to the provided param host
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

func (client *APIClient) getConversionURL(from, to currency.Currency, value decimal.Decimal) string {
	return fmt.Sprintf(client.apiHost+ConversionUri, from, to, value.String())
}
