package main

import (
	"fmt"

	"github.com/certainlyMohneeesh/golang-lets-fuckin-go/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("%v\n%v\n", s1, s2)
}
