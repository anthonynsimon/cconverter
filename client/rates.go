package client

import "fmt"

func (client *APIClient) getRatesURL(currency string) string {
	return fmt.Sprintf(client.apiHost+RatesUri, currency)
}

func (client *APIClient) getConversionURL(from, to, value string) string {
	return fmt.Sprintf(client.apiHost+ConversionUri, from, to, value)
}

func (client *APIClient) GetRates(currency string) (*ExchangeRates, error) {
	response, err := client.httpClient.Get(client.getRatesURL(currency))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	payload := apiResponseRates{}
	err = decodePayload(response.Body, &payload)
	if err != nil {
		return nil, err
	}

	err = validate(&payload)
	if err != nil {
		return nil, err
	}

	rates, err := payload.unfold()
	if err != nil {
		return nil, err
	}

	return rates, nil
}
