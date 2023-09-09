package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() // Get the current time, including monotonic clock reading

	// Simulate some work
	time.Sleep(2 * time.Second)

	end := time.Now() // Get the current time again

	elapsed := end.Sub(start) // Compute the elapsed time using monotonic clock

	fmt.Printf("Elapsed time: %s\n", elapsed)
}
