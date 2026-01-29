package puppy

import (
	"github.com/certainlyMohneeesh/golang-lets-fuckin-go/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func ManBark() string {
	return dog.WhenGrownUp(Bark())
}

func ManBarks() string {
	return dog.WhenGrownUp(Barks())
}
