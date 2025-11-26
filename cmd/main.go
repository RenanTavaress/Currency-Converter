package main

import (
	"currency-converter/internal/cli"
	"fmt"
	"os"
)

func main() {
	err := cli.GetInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n\n", err)
		os.Exit(1)
	}
}
