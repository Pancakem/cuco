package main

import (
	"net/http"
	"encoding/json"
	"strings"
	"strconv"
	"fmt"
)

type Response struct {
	Error string `json:"error"`
	Data interface{} `json:"data"`
}

func Handler() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/convert", ConvertHandler)
	router.HandleFunc("/conversion-table", ShowCurrencyConversionTableHandler)
	return router
}


// show conversion rates for one currency to the rest of the currencies
func ShowCurrencyConversionTableHandler(w http.ResponseWriter, r *http.Request) {
	// recover if there's a panic on indexing the url queries
	defer recoverFromPanic(w)

	currencies := []string{"ksh", "ghs", "ngn"}
	
	w.Header().Set("Content-Type", "application/json")
	var resp Response
	urlQueries := r.URL.Query()

	currency := urlQueries["currency"][0]
	if len(currency) == 0 || !IsSupported(currency) {
		resp.Error = "currency is either empty or not supported. NGN, KSH, GHS supported."
		resp.Data = nil
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&resp)
		return
	}

	outerdata := make(map[string]interface{})
	outerdata["base_currency"] = currency
	table := make(map[string]float64)	
	for _, v := range currencies {
		if v == currency {
			continue
		}
		table[v] = CalculateConversionRate(currency, v)
	}
	outerdata["rates"] = table
	resp.Error = "None"
	resp.Data = outerdata
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)
}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	// recover if there's a panic on indexing the url queries
	defer recoverFromPanic(w)
	
	w.Header().Set("Content-Type", "application/json")
	var resp Response
	urlQueries := r.URL.Query()
	
	to := urlQueries["to"][0]
	if len(to) == 0 || !IsSupported(to) {
		resp.Error = "to is either empty or not supported. NGN, KSH, GHS supported."
		resp.Data = nil
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&resp)
		return
	}

	from := urlQueries["from"][0]
	if len(from) == 0 || !IsSupported(from){
		resp.Error = "from is either empty or not supported. NGN, KSH, GHS supported."
		resp.Data = nil
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&resp)
		return
	}

	valueString := urlQueries["amount"][0]
	value, err := strconv.ParseFloat(valueString, 64)
	if  err != nil {
		resp.Error = "amount must be a number"
		resp.Data = nil
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&resp)
		return
	}

	converted := Convert(value, from, to)
	resp.Error = "None"
	resp.Data = struct {
		CurrencyCode string `json:"currency_code"`
		CurrencyValue float64 `json:"currency_value"`
	}{
		CurrencyCode: strings.ToUpper(to),
		CurrencyValue: converted,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)	
}


func recoverFromPanic(w http.ResponseWriter) {
	if r := recover(); r != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorString := fmt.Sprintf("Bad request: %v\nPlease supply values in the request\n", r)
		w.Write([]byte(errorString))
	}		
}
