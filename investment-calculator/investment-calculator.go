package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	years := 10.0
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	futurevalue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futurevalue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futurevalue)
	fmt.Println(futureRealValue)

}
