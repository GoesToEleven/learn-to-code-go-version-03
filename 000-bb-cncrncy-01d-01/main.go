package main

import (
	"fmt"
	"sync"
)

// problematic
// this is WITH closure
func main() {
	var wg sync.WaitGroup
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// anonymous wrapper func WITH closure; it references the current value of i
		go func() {
			fmt.Println("goroutine ", i)
			wg.Done()
		}()
	}
	
	wg.Wait()
	
	fmt.Println("i'm in goroutine 1 - main")
}
