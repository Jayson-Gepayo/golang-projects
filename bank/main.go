package main

import (
	"fmt"
	"golang-projects/bank/helpers"
)

const balanceFile = "balance.txt"

func main() {

	var choice int
	accountBalance, err := helpers.GetFloatFromFile(balanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------")
		panic("ERROR")
	}

	fmt.Println("WELCOME TO GOBANK")

	for {
		helpers.PresentOptions()
		fmt.Println("Your Choice")
		fmt.Scan(&choice)
		fmt.Println("Youv'e Choice: ", choice)

		switch choice {
		case 1:
			helpers.GetFloatFromFile(balanceFile)
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
			helpers.WriteFloatToFile(accountBalance, balanceFile)
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
			helpers.WriteFloatToFile(accountBalance, balanceFile)
		default:
			fmt.Println("Goodbye")
			fmt.Println("Thank you for Choosing our bank")
			return
		}
	}
}
