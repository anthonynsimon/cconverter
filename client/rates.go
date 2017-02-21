package client

import "github.com/anthonynsimon/cconverter/currency"

// GetRates returns an ExchangeRates struct holding the resulting data from the exchange rate query operation.
// Parameter 'currency' corresponds to the currency to be used as a base for the rates.
func (client *APIClient) GetRates(currency currency.Currency) (*ExchangeRates, error) {
	response, err := client.httpClient.Get(client.getRatesURL(currency))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if !isResponseStatusOK(response) {
		return nil, ErrResponseNotOk
	}

	rates := &ExchangeRates{}
	err = decodePayload(response.Body, rates)
	if err != nil {
		return nil, err
	}

	return rates, nil
}
