package client

import (
	"encoding/json"
	"io"
)

var (
	jsonContentType = "application/json; charset=utf-8"
)

type ExchangeRates struct {
	Rates map[string]float64 `json:"rates,omitempty"`
	Base  string             `json:"base,omitempty"`
}

func decodePayload(reader io.Reader, payload interface{}) error {
	err := json.NewDecoder(reader).Decode(&payload)
	if err != nil {
		return err
	}
	return nil
}
