package main

import "fmt"

// Курсы валют
const rateUSDEUR float64 = 0.9
const rateUSDRUB float64 = 80.0
const rateEURRUB float64 = rateUSDRUB / rateUSDEUR

func main() {
	fmt.Printf("Курс USD к RUB %.2f\n", rateUSDRUB)
	fmt.Printf("Курс USD к EUR %.2f\n", rateUSDEUR)
	fmt.Printf("Курс EUR к RUB %.2f\n", rateEURRUB)
}
