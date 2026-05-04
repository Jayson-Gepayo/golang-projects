package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Investment years: ")
	fmt.Scan(&years)

	futurevalue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futurevalue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value:", futurevalue)
	fmt.Println("Future Real value (adjusted for Inflation)", futureRealValue)

}
