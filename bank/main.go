package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {

	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFile)

	if err != nil {
		return 1000, errors.New("Failed to read file")
	}

	balanceData := string(data)
	formattedBalanceData, err := strconv.ParseFloat(balanceData, 64)

	if err != nil {
		return 1000, errors.New("failed to parse balance")
	}

	return formattedBalanceData, nil
}

func main() {

	var choice int
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------")
		panic("ERROR")
	}

	fmt.Println("WELCOME TO GOBANK")

	for {
		fmt.Println("What do you want to do")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. EXIT")

		fmt.Println("Your Choice")
		fmt.Scan(&choice)
		fmt.Println("Youv'e Choice: ", choice)

		switch choice {
		case 1:
			getBalanceFromFile()
			fmt.Println("your balance is: ", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("deposit amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("please enter a valid deposit amount")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			var widthdrawAmount float64
			fmt.Print("withdraw amount: ")
			fmt.Scan(&widthdrawAmount)

			if widthdrawAmount <= 0 {
				fmt.Println("please enter a valid deposit amount")
				continue
			}

			if widthdrawAmount > accountBalance {
				fmt.Println("inefficient amount")
				fmt.Println("your balance is:", accountBalance)
				continue
			}

			accountBalance -= widthdrawAmount
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye")
			fmt.Println("Thank you for Choosing our bank")
			return
		}
	}
}
