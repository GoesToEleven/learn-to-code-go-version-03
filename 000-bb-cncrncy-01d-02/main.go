package main

import (
	"fmt"
	"sync"
)

// solution #1
// PARAMETER: send the current value of i with the goroutine
func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// anonymous wrapper func WITHOUT closure
		go func(n int) {
			fmt.Println("goroutine ", n)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println("i'm in goroutine 1 - main")
}