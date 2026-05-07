package main

import (
	"errors"
	"fmt"
	"os"
)

const resultsFile = "results.txt"

func writeToFile(revenue, profit, ratio float64) {

	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.3f\n", revenue, profit, ratio)
	os.WriteFile(resultsFile, []byte(results), 0644)

}

func main() {

	var choice int

	fmt.Println("WELCOME TO PROFIT CALCULATOR")

	for {
		fmt.Println("What do you want to do")
		fmt.Println("1. Compute for Profit")
		fmt.Println("2. EXIT")
		fmt.Scan(&choice)

		if choice == 1 {
			revenue, err := getUserInput("Revenue: ")

			if err != nil {
				fmt.Println(err)
				continue
			}

			expenses, err := getUserInput("Expenses: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			taxRate, err := getUserInput("Tax Rate: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			ebt, profit, ratio := calculate(revenue, expenses, taxRate)

			fmt.Printf("Earnings Before Tax: %0.2f\n", ebt)
			fmt.Printf("Profit: %0.2f\n", profit)
			fmt.Printf("Ratio: %0.2f \n\n", ratio)
			writeToFile(ebt, profit, ratio)
		} else {
			fmt.Println("Goodbye")
			fmt.Println("Thank you")
			return
		}
	}
}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Value must be Positive")
	}
	return userInput, nil
}

func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
