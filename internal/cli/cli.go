package cli

import (
	"currency-converter/internal/usecase"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInput() error {
	args := os.Args[1:]

	if len(args) < 2 {
		return errors.New("error: insufficient arguments provided. Usage: <amount> <currency>")
	}

	coin, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return errors.New("error: invalid amount. The first argument must be a valid number")
	}
	currency := args[1]

	if currency == "" {
		return errors.New("error: currency code is required. Please provide a valid currency code (e.g., USD, BRL, EUR)")
	}

	currency = strings.ToUpper(currency)

	fee, err := usecase.ConvertCurrency(coin, currency)

	if err != nil {
		return err
	}

	fmt.Printf("%.2f\n", fee)

	return nil
}

// // Run processa os argumentos da linha de comando e retorna os valores
// func Run() (string, string, float64, error) {
// 	var (
// 		fromCurrency = flag.String("from", "", "Moeda de origem (ex: USD)")
// 		toCurrency   = flag.String("to", "", "Moeda de destino (ex: BRL)")
// 		amount       = flag.Float64("amount", 0, "Valor a ser convertido")
// 	)

// 	flag.Parse()

// 	// Validação dos argumentos obrigatórios
// 	if *fromCurrency == "" {
// 		return "", "", 0, fmt.Errorf("erro: flag -from é obrigatória")
// 	}
// 	if *toCurrency == "" {
// 		return "", "", 0, fmt.Errorf("erro: flag -to é obrigatória")
// 	}
// 	if *amount <= 0 {
// 		return "", "", 0, fmt.Errorf("erro: flag -amount deve ser maior que zero")
// 	}

// 	return *fromCurrency, *toCurrency, *amount, nil
// }

// // ShowUsage exibe a mensagem de uso do programa
// func ShowUsage() {
// 	fmt.Fprintf(os.Stderr, "Uso: %s [opções]\n\n", os.Args[0])
// 	fmt.Fprintf(os.Stderr, "Opções:\n")
// 	flag.PrintDefaults()
// }
