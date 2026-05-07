package main

import "fmt"

func main() {

	var choice int
	var accountBalance float64 = 10000
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

		if choice == 1 {
			fmt.Println("your balance is: ", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("deposit amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("please enter a valid deposit amount")
				continue
			}

			accountBalance += depositAmount

			fmt.Println("Balance updated! New amount:", accountBalance)

		} else if choice == 3 {
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

		} else {
			fmt.Println("Goodbye")
			break
		}
	}
	fmt.Println("Thank you for Choosing our bank")
}

// func checkBalance() {

// }
