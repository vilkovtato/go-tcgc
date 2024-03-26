package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 0, err
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0, err
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("Error reading balance from file: ", err)
		return
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Println("Enter amount to deposit: ")
			var amount float64
			fmt.Scan(&amount)
			accountBalance += amount
			fmt.Println("Your new balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Println("Enter amount to withdraw: ")
			var amount float64
			fmt.Scan(&amount)

			if amount > accountBalance {
				fmt.Println("You don't have enough money!")
				continue
			}

			if amount <= 0 {
				fmt.Println("Invalid amount!")
				continue
			}

			accountBalance -= amount
			fmt.Println("Your new balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Bye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}
