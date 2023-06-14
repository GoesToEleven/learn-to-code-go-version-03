package main

import (
	"fmt"
	"time"
)

func main() {

	// we won't see this print
	go fmt.Println("goroutine 2")
	fmt.Println("i'm in goroutine 1 - main")
	time.Sleep(10 * time.Millisecond)
}

// The go runtime does not wait for go routines
