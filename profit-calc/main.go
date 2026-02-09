package main

import (
	"fmt"
	"strings"
	"time"
)

//Goals
// 1) Validate user input
// 		=> Show error message & exit if invalid input is provided
// 		- No negative numbers
// 		- Not 0
// 2) Store calculated results into file

// var revenue float64
// var expenses float64
// var taxRate float64
func outputText(text string) {
	fmt.Print(text)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate)
	return ebt, profit
}

func getUserInput(infoText string) float64 {
	var userInput float64
	outputText(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		fmt.Println("Invalid or negative input.\nEnter a valid input once again.")
		return getUserInput(infoText)
	}

	return userInput
}

func main() {

	outputText("╔══════════════════════════════════════════════════════════════════════════════╗\n")
	outputText("║                          Welcome to the Profit Calculator                    ║\n")
	outputText("╚══════════════════════════════════════════════════════════════════════════════╝\n")
	for {

		outputText("1. Start Calculating!!\n")
		outputText("2. Exit\n")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			revenue := getUserInput("Enter Revenue: ")
			expenses := getUserInput("Enter Expenses: ")
			taxRate := getUserInput("Enter taxRate: ")
			taxRate = taxRate / 100

			// just my fancy loading text
			outputText("Calculating profit")
			for i := 0; i < 3; i++ {
				time.Sleep(500 * time.Millisecond)
				outputText(".")
			}
			//this clears the calculating loading text
			outputText("\r" + strings.Repeat(" ", len("Calculating profit...")) + "\r")

			// ebt := revenue - expenses
			// ebt := calculateFinancials(revenue - expenses)
			// profit := ebt * (1 - taxRate)
			// profit := calculateFinancials(ebt * (1 - taxRate))

			ebt, profit := calculateFinancials(revenue, expenses, taxRate)

			fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
			fmt.Printf("Net Profit: %.2f\n", profit)
			fmt.Printf("Tax Burden Ratio: %.2f\n", ebt/profit)
			fmt.Printf("Tax Retention Rate: %.2f\n", profit/ebt)
			fmt.Printf("Pretax Margin: %.2f\n", ebt/revenue)
			fmt.Printf("Net Profit Margin: %.2f\n", profit/revenue)

		default:
			fmt.Println("Goodbye!\nWill see you again later.")
			return
		}

	}
}
