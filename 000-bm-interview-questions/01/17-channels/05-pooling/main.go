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
	pooling()
}

// pooling: In this pattern, the parent goroutine signals 100 pieces of work
// to a pool of child goroutines waiting for work to perform.
func pooling() {

	// GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously
	// an argument of 0 returns the current number of logical CPUs used by the Go runtime.
	g := runtime.GOMAXPROCS(0)

	// NumCPU returns the number of logical CPUs usable by the current process.
	// This is the total number of logical CPUs on the machine, regardless of the current GOMAXPROCS setting.
	// This value can be useful when you're trying to make decisions about the level of concurrency in your program
	// or understand the hardware capabilities of the machine where the process is running.
	fmt.Println("num CPU's:", runtime.NumCPU())

	ch := make(chan string)

	for c := 0; c < g; c++ {
		go func(child int) {
			for d := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", child, d)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(c)
	}

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "data"
		fmt.Println("parent : sent signal :", w)
	}

	close(ch)
	fmt.Println("parent : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
