package currency

import (
	"errors"
	"strings"
)

type Currency string

const (
	AUD Currency = "AUD"
	BGN Currency = "BGN"
	BRL Currency = "BRL"
	CAD Currency = "CAD"
	CHF Currency = "CHF"
	CNY Currency = "CNY"
	CZK Currency = "CZK"
	DKK Currency = "DKK"
	GBP Currency = "GBP"
	HKD Currency = "HKD"
	HRK Currency = "HRK"
	HUF Currency = "HUF"
	IDR Currency = "IDR"
	ILS Currency = "ILS"
	INR Currency = "INR"
	JPY Currency = "JPY"
	KRW Currency = "KRW"
	MXN Currency = "MXN"
	MYR Currency = "MYR"
	NOK Currency = "NOK"
	NZD Currency = "NZD"
	USD Currency = "USD"
	PHP Currency = "PHP"
	PLN Currency = "PLN"
	RON Currency = "RON"
	RUB Currency = "RUB"
	SEK Currency = "SEK"
	SGD Currency = "SGD"
	THB Currency = "THB"
	TRY Currency = "TRY"
	ZAR Currency = "ZAR"
	EUR Currency = "EUR"
)

var knownCurrencies = []Currency{
	"AUD",
	"BGN",
	"BRL",
	"CAD",
	"CHF",
	"CNY",
	"CZK",
	"DKK",
	"GBP",
	"HKD",
	"HRK",
	"HUF",
	"IDR",
	"ILS",
	"INR",
	"JPY",
	"KRW",
	"MXN",
	"MYR",
	"NOK",
	"NZD",
	"USD",
	"PHP",
	"PLN",
	"RON",
	"RUB",
	"SEK",
	"SGD",
	"THB",
	"TRY",
	"ZAR",
	"EUR",
}

func Parse(str string) (Currency, error) {
	if len(str) != 3 {
		return "", errors.New("invalid currency code length")
	}
	str = strings.ToUpper(str)
	parsed := Currency(str)
	for _, currency := range knownCurrencies {
		if parsed == currency {
			return currency, nil
		}
	}
	return "", errors.New("unkown currency code")
}
