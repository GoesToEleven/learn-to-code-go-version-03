// go build -race
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		atomic.AddInt64(&counter, 1)
		wg.Done()
	}()

	go func() {
		atomic.AddInt64(&counter, 1)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
