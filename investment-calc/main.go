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

	outputText("Enter Your Investment Amount: ")
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future Value: %v\n", futureValue)
	fmt.Printf("Future Value (adjusted for Inflation):  %v\n", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}
