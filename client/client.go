package client

import (
	"net/http"
	"time"
)

var (
	APIRequestTimeout = time.Second * 10
)

const (
	RatesUri      = "/api/rates/%s"
	ConversionUri = "/api/convert?from=%s&to=%s&amount=%s"
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
