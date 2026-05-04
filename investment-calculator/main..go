package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Investment years: ")
	outputText("Investment years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := futureValueCalculate(investmentAmount, expectedReturnRate, years)
	// fmt.Println("Future value:", futurevalue)
	// fmt.Printf("Future value: %.2f\nFuture Real value (adjusted for Inflation) %.2f\n", futurevalue, futureRealValue)
	// fmt.Println("Future Real value (adjusted for Inflation)", futureRealValue)

	formatted := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real value (adjusted for Inflation) %.2f\n", futureRealValue)

	fmt.Print(formatted, formattedRFV)

}

func outputText(text string) {
	fmt.Print(text)
}

func futureValueCalculate(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
