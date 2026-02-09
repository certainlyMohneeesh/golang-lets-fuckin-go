package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance = getBalanceFromFile()

	fmt.Println("Welcome to THEBANK!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println(`1. Check Balance
2. Deposit Money
3. Withdraw money
4. Exit`)

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1
		switch choice {
		case 1:
			writeBalanceToFile(accountBalance)
			fmt.Printf("Your balance is: %.2f\n", accountBalance)

		case 2:
			fmt.Print("Enter the Deposit Amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Printf("Your balance updated!\nCurrent Balance: %.2f\n", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Enter the Withdrawal Amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount.\nMust be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Sorry!\nInsufficient Balance!!!")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Printf("Your balance updated!\nCurrent Balance: %.2f\n", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Thank you for choosing our bank!\nGoodbye!!")
			return
		}
	}

}
