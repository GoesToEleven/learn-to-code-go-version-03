package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	drop()
}

// drop: In this pattern, the manager goroutine signals 2000 pieces of work to
// a single worker goroutine that can't handle all the work. If the manager
// performs a send and the worker is not ready, that work is discarded and dropped.
func drop() {
	const cap = 100
	ch := make(chan int, cap)

	go func() {
		for p := range ch {
			fmt.Println("worker rcvd signal", p)
		}
	}()

	const work = 2000
	var dropped int
	for i := 0; i < work; i++ {
		select {
		case ch <- i:
			fmt.Println("mgr sent signal", i)
		default:
			fmt.Println("mgr dropped data", i)
			dropped++
		}
	}

	close(ch)
	fmt.Println("mgr sent shutdown signal")
	fmt.Println("dropped", dropped)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
