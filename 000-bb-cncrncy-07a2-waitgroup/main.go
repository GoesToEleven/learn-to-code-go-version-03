package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go hello(&wg)
	wg.Wait()
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello")
}
