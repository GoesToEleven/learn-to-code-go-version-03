// go build -race
package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		mu.Lock()
		{
			counter++
		}
		mu.Unlock()
		wg.Done()
	}()

	go func() {
		mu.Lock()
		{
			counter++
		}
		mu.Unlock()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
