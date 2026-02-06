package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	outputText("-------------------Welcome to the Profit Calculator---------------------\n")

	outputText("Enter Revenue: ")
	fmt.Scan(&revenue)
	outputText("Enter Expenses: ")
	fmt.Scan(&expenses)
	outputText("Enter the Tax Rate: ")
	fmt.Scan(&taxRate)

	taxRate = taxRate / 100

	// just my fancy loading text
	outputText("Calculating profit")
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		outputText(".")
	}
	//this clears the calculating loading text
	outputText("\r" + strings.Repeat(" ", len("Calculating profit...")) + "\r")

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate)

	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
	fmt.Printf("Net Profit: %.2f\n", profit)
	fmt.Printf("Tax Burden Ratio: %.2f\n", ebt/profit)
	fmt.Printf("Tax Retention Rate: %.2f\n", profit/ebt)
	fmt.Printf("Pretax Margin: %.2f\n", ebt/revenue)
	fmt.Printf("Net Profit Margin: %.2f\n", profit/revenue)
}

func outputText(text string) {
	fmt.Print(text)
}
