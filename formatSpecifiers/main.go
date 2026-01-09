package main

import "fmt"

func main() {
	i := 42

	fmt.Printf("Decimal: %d\n", i)
	fmt.Printf("Binary: %b\n", i)
	fmt.Printf("Default: %v\n", i)
	fmt.Printf("Hexadecimal (uppercase): %X\n", i)
	fmt.Printf("Hexadecimal (lowercase): %x\n", i)
}
