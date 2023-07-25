package main

import (
	"fmt"
	"math/rand"
	"time"
)

type car struct{}

// String implements the fmt.Stringer interface.
func (car) String() string {
	return "Vroom!"
}

type cloud struct{}

// String implements the fmt.Stringer interface.
func (cloud) String() string {
	return "Big Data!"
}

func main() {

	// Seed the number random generator.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a slice of the Stringer interface values.
	mvs := []fmt.Stringer{
		car{},
		cloud{},
	}

	// Let's run this experiment ten times.
	for i := 0; i < 10; i++ {

		// Choose a random number from 0 to 1.
		rn := r.Intn(2)

		// Perform a type assertion that we have a concrete type
		// of cloud in the interface value we randomly chose.
		if v, ok := mvs[rn].(cloud); ok {
			fmt.Println("Got Lucky:", v)
			continue
		}

		fmt.Println("Got Unlucky")
	}
}
