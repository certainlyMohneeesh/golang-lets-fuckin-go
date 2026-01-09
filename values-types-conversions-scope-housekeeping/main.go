package main

import "fmt"

func main() {
	// Variable declaration and initialization
	a := 10.5
	b := 20.5
	c := "Hello, Go!"

	fmt.Printf("%T: %v\n", a, a)
	fmt.Printf("%T: %v\n", b, b)
	fmt.Printf("%T: %v\n", c, c)

	fmt.Println("Concatenated string:", c+" Let's learn Go!")

	fmt.Printf("Sum of float and float (after conversion): %v\n", a+b)
}
