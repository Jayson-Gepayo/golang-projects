package main

import (
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {

	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(balanceFile)
	balanceData := string(data)
	formattedBalanceData, _ := strconv.ParseFloat(balanceData, 64)

	return formattedBalanceData
}

func main() {

	var choice int
	var accountBalance float64 = getBalanceFromFile()
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
