package main

import (
	"fmt"
	"strings"
)

// values of x currency to ksh
// ksh is our base currency
// having a base currency let's us have a smaller table of currencies 
var currencyExchangeTable = map[string]float64{
	"ngn": 3.8040,
	"ghs": 0.055040,
}

// get the currency from value if none is supplied
func ParseCurrencyCode(valueString string) string {
	code := ""
	for _, v := range valueString {
		// skip over digits
		if v >= 48 && v <= 57 {
			continue
		}

		code = fmt.Sprintf("%s%c", code, v)
	}

	return code
}

// checks if the supplied currency is supported
func IsSupported(code string) bool {
	res := false
	switch strings.ToLower(code) {
	case "ksh", "ngn", "ghs":
		res = true
	default:
		res = false
	}
	return res
}

func Convert(value float64, from, to string) float64 {
	return toTwoDecimalPlaces(CalculateConversionRate(from, to) * value)
}

// computes the appropriate conversion rates
func CalculateConversionRate(from, to string) float64 {
	if strings.ToLower(from) == "ksh" {
		return currencyExchangeTable[to]
	}

	if strings.ToLower(to) == "ksh" {
		return 1 / currencyExchangeTable[from]
	}

	return currencyExchangeTable[to] / currencyExchangeTable[from]
}
