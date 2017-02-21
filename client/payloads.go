package client

import (
	"encoding/json"
	"io"
)

// ExchangeRates corresponds to the payload returned when querying exchange rates
// for a base currency.
type ExchangeRates struct {
	Rates map[string]float64 `json:"rates,omitempty"`
	Base  string             `json:"base,omitempty"`
}

// ExchangeQuote corresponds to the payload returned when converting an amount
// between currencies.
type ExchangeQuote struct {
	FromCurrency     string  `json:"fromCurrency,omitempty"`
	ToCurrency       string  `json:"toCurrency,omitempty"`
	AmountToConvert  float64 `json:"amountToConvert,omitempty"`
	ConversionResult float64 `json:"conversionResult,omitempty"`
	ExchangeRate     float64 `json:"exchangeRate,omitempty"`
}

func decodePayload(reader io.Reader, payload interface{}) error {
	err := json.NewDecoder(reader).Decode(&payload)
	if err != nil {
		return err
	}
	return nil
}
