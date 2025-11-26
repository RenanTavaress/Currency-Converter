package usecase

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type CurrencyConverter struct {
	BaseCurrency string             `json:"base"`
	Rates        map[string]float64 `json:"rates"`
	Date         string             `json:"date"`
}

func ConvertCurrency(coin float64, currency string) (float64, error) {

	var currencyData CurrencyConverter
	jsonFile, err := os.Open("currency.json")
	if err != nil {
		return 0, err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	if err := json.Unmarshal(byteValue, &currencyData); err != nil {
		return 0, errors.New("error: failed to parse currency data from JSON file")
	}

	fee, ok := currencyData.Rates[currency]

	if !ok {
		return 0, errors.New("error: currency not found. Please provide a valid currency code")
	}

	fee = fee * coin

	return fee, nil
}
