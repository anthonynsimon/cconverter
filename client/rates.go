package client

import (
	"fmt"

	"github.com/anthonynsimon/xeclient/currency"
)

func (client *APIClient) getRatesURL(currency currency.Currency) string {
	return fmt.Sprintf(client.apiHost+RatesUri, currency)
}

func (client *APIClient) getConversionURL(from, to currency.Currency, value string) string {
	return fmt.Sprintf(client.apiHost+ConversionUri, from, to, value)
}

func (client *APIClient) GetRates(currency currency.Currency) (*ExchangeRates, error) {
	response, err := client.httpClient.Get(client.getRatesURL(currency))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	rates := &ExchangeRates{}
	err = decodePayload(response.Body, rates)
	if err != nil {
		return nil, err
	}

	return rates, nil
}
