package main

import (
	"fmt"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("The addition of two random integer numbers are: %v\n", add(rand.Intn(50), rand.Intn(50)))
}
