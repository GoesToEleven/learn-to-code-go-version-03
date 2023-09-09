package main

import (
	"fmt"
	"time"
)

func expensiveCall() {
	time.Sleep(10 * time.Millisecond)
}

func main() {
	t0 := time.Now()
	expensiveCall()

	// one way
	t1 := time.Now()
	fmt.Printf("The call took %v ms to run.\n", t1.Sub(t0))

	// second way
	fmt.Printf("The call took %v ms to run.\n", time.Now().Sub(t0))

	// third way since
	fmt.Printf("The call took %v ms to run.\n", time.Since(t0))

}
