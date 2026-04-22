package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalance = "account_balance.txt"

func writeBalanceToFile(balance float64) {
	balaceText := fmt.Sprint(balance)
	os.WriteFile(accountBalance, []byte(balaceText), 0644)
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalance)

	if err != nil {
		return 1000, errors.New("FAILED TO READ FILE")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("FAILED TO PARSE STORED BALANCE VALUE")
	}
	return balance, nil
}

func main() {

	var accountBalance, err = readBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------")
		panic(err)
	}
	var choice int
	fmt.Println("WELCOME TO GO BANK")

	for {
		fmt.Println("What do yo want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance: ", accountBalance)
		case 2:
			fmt.Print("Deposit Amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("INVALID AMOUNT")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("New Balance Amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Withdraw Amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("INVALID AMOUNT WITHDRAW AMOUNT")
				continue
			} else if withdrawAmount >= accountBalance {
				fmt.Println("insufficient amount")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("New Balance Amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye")
			fmt.Println("GOODBYE 2")
			return
		}
	}

}
