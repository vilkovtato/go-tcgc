package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Enter number of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFinancials(investmentAmount, expectedReturnRate, years)
	
	formatedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formatedFVA := fmt.Sprintf("Future value (adjusted for inflation): %.1f\n", futureRealValue)

	fmt.Print(formatedFV, formatedFVA)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFinancials(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {

	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
    rfv = fv / math.Pow(1 + inflationRate / 100, years)

	return fv, rfv
}