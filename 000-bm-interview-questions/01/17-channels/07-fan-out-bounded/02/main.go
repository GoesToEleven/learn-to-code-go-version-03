package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	boundedWorkPooling()
}

// boundedWorkPooling: In this pattern, a pool of worker goroutines is created
// to service a fixed amount of work. The parent goroutine iterates over all
// work, signalling that into the pool. Once all the work has been signaled,
// then the channel is closed, the channel is flushed, and the worker
// goroutines terminate.

func boundedWorkPooling() {

	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	start := time.Now()
	for i := 0; i < g; i++ {
		go func(worker int) {
			defer wg.Done()
			for wrk := range ch {
				fmt.Printf("rcvd signal by worker %d - %s\n", worker, wrk)
			}
			fmt.Printf("rcvd shutdown signal by worker %d\n", worker)
		}(i)
	}

	// sets index position 20 for last item
	work := []string{"paper", "paper", "paper", "paper", 20: "paper"}
	for _, w := range work {
		ch <- w
	}
	close(ch)
	wg.Wait()

	fmt.Printf("%s \n", time.Since(start))
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
