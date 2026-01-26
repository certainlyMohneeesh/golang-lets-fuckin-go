package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to KBC!!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")

	// comma ok pattern

	input, _ := reader.ReadString('\n')
	fmt.Println("Hooray!! you won Saath Crorez", input)
	fmt.Printf("Type of the input is %T", input)
}
