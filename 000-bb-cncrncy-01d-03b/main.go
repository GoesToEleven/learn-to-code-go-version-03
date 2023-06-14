package main

import (
	"fmt"
	"sync"
)

// solution #3
// VARIABLE SHADOWING: send the current value of i with the goroutine
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		j := i // more readable
		go func() {
			fmt.Println("goroutine ", j)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("i'm in goroutine 1 - main")
}
