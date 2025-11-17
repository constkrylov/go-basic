package main

import "fmt"

// Курсы валют
const rateUSDEUR float64 = 0.9
const rateUSDRUB float64 = 80.0
const rateEURRUB float64 = rateUSDRUB / rateUSDEUR

func getUserInput() (fromCurrency string, toCurrency string, amount float64) {
	fmt.Println("Введите исходную валюту:")
	fmt.Scanln(&fromCurrency)
	fmt.Println("Введите целевую валюту:")
	fmt.Scanln(&toCurrency)
	fmt.Println("Введите сумму:")
	fmt.Scanln(&amount)
	return fromCurrency, toCurrency, amount
}

func convertCurrency(fromCurrency string, toCurrency string, amount float64) float64 {
	return amount
}

func main() {
	fmt.Printf("Курс USD к RUB %.2f\n", rateUSDRUB)
	fmt.Printf("Курс USD к EUR %.2f\n", rateUSDEUR)
	fmt.Printf("Курс EUR к RUB %.2f\n", rateEURRUB)
}
