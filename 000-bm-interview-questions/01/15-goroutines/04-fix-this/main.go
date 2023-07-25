package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(2)
}

// fix this code
func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		fmt.Println("process one")
		wg.Done()
	}()

	go func() {
		fmt.Println("process two")
		wg.Done()
	}()

	wg.Wait()
}
