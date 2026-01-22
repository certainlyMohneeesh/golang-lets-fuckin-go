package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var play bool

	fmt.Print("Do you want to play the dice game? (true/false): ")

	_, err := fmt.Scan(&play)

	if err != nil {
		fmt.Println("Invalid input. Please enter true or false.")
		return
	}

	if play {

		diceResult := rand.Intn(6) + 1

		fmt.Printf("Rolling the dice...\n")
		time.Sleep(500 * time.Millisecond)

		fmt.Printf("You rolled a: %d\n", diceResult)
	} else {
		fmt.Println("Maybe next time!")
	}
}
