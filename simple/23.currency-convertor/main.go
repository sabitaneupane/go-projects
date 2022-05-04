package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	EUR = iota
	GBP
	JPY
)

var rates = []float64{
	EUR: 0.88,
	GBP: 0.78,
	JPY: 113.02,
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please input amount.")
		return
	}

	amount, err := strconv.ParseFloat(os.Args[1], 64)

	if err != nil {
		fmt.Println("Invalid amount. It should be a number")
	}

	fmt.Printf("%.2f USD is %.2f EUR\n", amount, rates[EUR]*amount)
	fmt.Printf("%.2f USD is %.2f GBP\n", amount, rates[GBP]*amount)
	fmt.Printf("%.2f USD is %.2f JPY\n", amount, rates[JPY]*amount)
}
