package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	fanOutSem()
}

// fanOutSem: In this pattern, a semaphore is added to the fan out pattern
// to restrict the number of child goroutines that can be schedule to run.
func fanOutSem() {
	workers := 10
	ch := make(chan int)

	g := runtime.GOMAXPROCS(0)
	sem := make(chan bool, g)

	start := time.Now()
	for i := 0; i < workers; i++ {
		go func(worker int) {
			// when the buffered channel is full, this will block
			// this guarantees only runtime.GOMAXPROCS(0) goroutines are working at a time
			sem <- true
			{
				t := time.Duration(r.Intn(200)) * time.Millisecond
				time.Sleep(t)
				ch <- worker
				fmt.Println("signal sent by worker:", worker)
			}
			<-sem
		}(i)
	}

	for workers > 0 {
		d := <-ch
		workers--
		fmt.Println(d)
		fmt.Println("signal rcvd by parent:", workers)
	}
	fmt.Println(time.Since(start))

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
