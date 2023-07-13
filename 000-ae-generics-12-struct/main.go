package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type customData interface {
	constraints.Ordered | []byte | []rune
}

type User[T customData] struct {
	ID   int
	Name string
	Data T
}

func main() {
	u1 := User[int]{
		ID:   7,
		Name: "James",
		Data: 32,
	}
	fmt.Printf("%+v \n", u1)

	u2 := User[string]{
		ID:   8,
		Name: "Jenny",
		Data: "Always Paris",
	}
	fmt.Printf("%+v \n", u2)

}

// thank you to twitter @Fryancodez for the inspiration on this!
