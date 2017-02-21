package client

import "github.com/anthonynsimon/cconverter/currency"
import "github.com/shopspring/decimal"

// Convert returns an ExchangeQuote struct holding the resulting data from the conversion operation.
// Parameter 'from' corresponds to the base currency (the currency of the provided amount).
// Parameter 'to' corresponds to the target currency.
// Parameter 'amount' corresponds to the actual numerical amount to be converted.
func (client *APIClient) Convert(from, to currency.Currency, amount decimal.Decimal) (*ExchangeQuote, error) {
	convertURL := client.getConversionURL(from, to, amount)
	response, err := client.httpClient.Get(convertURL)
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
