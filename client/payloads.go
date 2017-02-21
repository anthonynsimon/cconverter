package client

import (
	"encoding/json"
	"io"
)

// TODO: set accept headers when making requests to the API
// var (
// 	jsonContentType = "application/json; charset=utf-8"
// )

// TODO: use a type better suited for financial numbers
type ExchangeRates struct {
	Rates map[string]float64 `json:"rates,omitempty"`
	Base  string             `json:"base,omitempty"`
}

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
