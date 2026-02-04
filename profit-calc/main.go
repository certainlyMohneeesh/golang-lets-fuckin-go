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

	fmt.Print("-------------------Welcome to the Profit Calculator---------------------\n")

	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the Tax Rate: ")
	fmt.Scan(&taxRate)

	taxRate = taxRate / 100

	// just my fancy loading text
	fmt.Print("Calculating profit")
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
	//this clears the calculating loading text
	fmt.Print("\r" + strings.Repeat(" ", len("Calculating profit...")) + "\r")

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate)

	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
	fmt.Printf("Net Profit: %.2f\n", profit)
	fmt.Printf("Tax Burden Ratio: %.2f\n", ebt/profit)
	fmt.Printf("Tax Retention Rate: %.2f\n", profit/ebt)
	fmt.Printf("Pretax Margin: %.2f\n", ebt/revenue)
	fmt.Printf("Net Profit Margin: %.2f\n", profit/revenue)
}
