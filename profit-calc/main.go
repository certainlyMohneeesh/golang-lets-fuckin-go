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

	fmt.Print("Enter Your Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
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

	fmt.Printf("Your Earning before tax is: %.2f\n", ebt)
	fmt.Printf("Your total Profit is: %.2f\n", profit)
}
