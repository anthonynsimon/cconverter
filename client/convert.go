package client

import "github.com/anthonynsimon/cconverter/currency"

func (client *APIClient) Convert(from, to currency.Currency, amount float64) (*ExchangeQuote, error) {
	convertUrl := client.getConversionURL(from, to, amount)
	response, err := client.httpClient.Get(convertUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if !isResponseStatusOK(response) {
		return nil, ErrResponseNotOk
	}

	quote := &ExchangeQuote{}
	err = decodePayload(response.Body, quote)
	if err != nil {
		return nil, err
	}

	return quote, nil
}
