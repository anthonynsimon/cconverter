package main

import (
	"encoding/json"
	"io"
	"time"
)

var (
	jsonContentType = "application/json; charset=utf-8"
)

type ValidatablePayload interface {
	Validate() error
}

type ExchangeRates struct {
	Rates map[string]float64
	Base  string
	Date  time.Time
}

type apiResponseRates struct {
	Rates map[string]float64 `json:"rates,omitempty"`
	Base  string             `json:"base,omitempty"`
	Date  string             `json:"date,omitempty"`
}

func (p *apiResponseRates) Validate() error {
	// TODO: validate
	return nil
}

func decodePayload(reader io.Reader, payload interface{}) error {
	err := json.NewDecoder(reader).Decode(&payload)
	if err != nil {
		return err
	}
	return nil
}

func validate(payload ValidatablePayload) error {
	if err := payload.Validate(); err != nil {
		return err
	}
	return nil
}

func (resp *apiResponseRates) unfold() (*ExchangeRates, error) {
	result := &ExchangeRates{}
	result.Base = resp.Base
	result.Rates = resp.Rates

	t, err := time.ParseInLocation("2006-01-02", resp.Date, time.FixedZone("", 0))
	if err != nil {
		return nil, err
	}
	result.Date = t
	return result, nil
}
