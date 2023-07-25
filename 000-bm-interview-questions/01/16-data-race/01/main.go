// go build -race
package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			value := counter
			value++
			counter = value
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			value := counter
			value++
			counter = value
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
