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
		i := i // variable shadowing
		go func() {
			fmt.Println("goroutine ", i)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("i'm in goroutine 1 - main")
}
