package main

import (
	"fmt"
	"os"
)

const revenueFile = "revenue.txt"
const expensesFile = "expenses.txt"
const taxRateFile = "taxRate.txt"
const ratioFile = "ratio.txt"

func writeToFile(revenue, expenses, tax, ratio float64) {

	revenueText := fmt.Sprint(revenue)
	os.WriteFile(revenueFile, []byte(revenueText), 0644)

	expensesText := fmt.Sprint(expenses)
	os.WriteFile(expensesFile, []byte(expensesText), 0644)

	taxRate := fmt.Sprint(tax)
	os.WriteFile(taxRateFile, []byte(taxRate), 0644)

	ratioText := fmt.Sprint(ratio)
	os.WriteFile(ratioFile, []byte(ratioText), 0644)
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
			revenue := getUserInput("Revenue: ")
			if revenue <= 0 {
				fmt.Println("Please enter valid revenue")
				continue
			}
			expenses := getUserInput("Expenses: ")
			if expenses <= 0 {
				fmt.Println("Please valid expenses")
				continue
			}
			taxRate := getUserInput("Tax Rate: ")
			if taxRate <= 0 {
				fmt.Println("Please valid tax rate")
				continue
			}

			ebt, profit, ratio := calculate(revenue, expenses, taxRate)

			fmt.Printf("Earnings Before Tax: %0.2f\n", ebt)
			fmt.Printf("Profit: %0.2f\n", profit)
			fmt.Printf("Ratio: %0.2f \n\n", ratio)
			writeToFile(ebt, profit, taxRate, ratio)
		} else {
			fmt.Println("Goodbye")
			fmt.Println("Thank you")
			return
		}
	}
}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
