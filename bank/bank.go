package main

import "fmt"

func main() {
	var accountBalance = 1000.00
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

		if choice == 1 {
			fmt.Printf("Your balance is: %.2f\n", accountBalance)
		} else if choice == 2 {
			fmt.Print("Enter the Deposit Amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Printf("Your balance updated!\nCurrent Balance: %.2f\n", accountBalance)
		} else if choice == 3 {
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
		} else if choice > 4 {
			fmt.Println("Enter a valid choice!\nYou must continue.")
			continue
		} else {
			fmt.Println("Thank you for choosing our bank!\nGoodbye!!")
			break
		}
	}

}
