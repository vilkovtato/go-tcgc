package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err := getUserInput("Enter revenue: ")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	expenses, err := getUserInput("Enter expenses: ")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	taxRate, err := getUserInput("Enter tax rate: ")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	ebt, profit, ratio := CaclucateValues(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.1f\n", ratio)
	fmt.Printf("%.3f\n", ebt)
}

func getUserInput(prompt string) (float64, error) {
	var value float64
	fmt.Print(prompt)
	fmt.Scan(&value)

	if value <= 0 {
		return 0, errors.New("input must be greater than 0")
	}

	return value, nil
}

func CaclucateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * ((100 - taxRate) / 100)
	ratio := ebt / profit

	writeBalanceToFile(ebt, profit, ratio)

	return ebt, profit, ratio
}

func writeBalanceToFile(ebt, profit, ratio float64) {

	text := fmt.Sprintf("profit: %.1f\nratio: %.1f\nEbt: %.3f\n", profit, ratio, ebt)

	os.WriteFile("balance.txt", []byte(text), 0644)
}
