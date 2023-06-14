package main

import (
	"fmt"
	"sync"
)

func main() {

	count := 0
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10_000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

// https://go.dev/doc/articles/race_detector#Runtime_Overheads