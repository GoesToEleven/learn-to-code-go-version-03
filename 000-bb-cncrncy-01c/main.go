package main

import (
	"fmt"
	"sync"
)

func main() {

	// sync.Waitgroup
	var wg sync.WaitGroup
	wg.Add(1)

	// anonymous wrapper func
	go func() {
		fmt.Println("goroutine 2")
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("i'm in goroutine 1 - main")
}
		