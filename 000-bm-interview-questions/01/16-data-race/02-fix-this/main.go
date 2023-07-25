// go build -race
package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		counter++
		wg.Done()
	}()

	go func() {
		counter++
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
