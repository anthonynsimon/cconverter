package client

import (
	"encoding/json"
	"io"

	"github.com/shopspring/decimal"
)

// ExchangeRates corresponds to the payload returned when querying exchange rates
// for a base currency.
type ExchangeRates struct {
	Rates map[string]decimal.Decimal `json:"rates,omitempty"`
	Base  string                     `json:"base,omitempty"`
}

// ExchangeQuote corresponds to the payload returned when converting an amount
// between currencies.
type ExchangeQuote struct {
	FromCurrency     string          `json:"fromCurrency,omitempty"`
	ToCurrency       string          `json:"toCurrency,omitempty"`
	AmountToConvert  decimal.Decimal `json:"amountToConvert,omitempty"`
	ConversionResult decimal.Decimal `json:"conversionResult,omitempty"`
	ExchangeRate     decimal.Decimal `json:"exchangeRate,omitempty"`
}

func decodePayload(reader io.Reader, payload interface{}) error {
	err := json.NewDecoder(reader).Decode(&payload)
	if err != nil {
		return err
	}
	return nil
}
