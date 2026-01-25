package main

import "fmt"

func main() {
	// Variable declaration and initialization
	a := 10.5
	b := 289.5
	c := "Hello, Go!"
	isLoggedIn := true

	fmt.Printf("Variable `a` is of type %T and its value is %v\n", a, a)
	fmt.Printf("Variable `b` is of type %T and its value is %v\n", b, b)
	fmt.Printf("Variable `c` is of type %T and its value is %v\n", c, c)
	fmt.Printf("The user is logged in: %v\n", isLoggedIn)

	fmt.Println("Concatenated string:", c+" Let's learn Go!")

	fmt.Printf("Sum of float and float (after conversion): %v\n", a+b)
}
