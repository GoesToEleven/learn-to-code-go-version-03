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

// pooling: In this pattern, the parent goroutine signals N pieces of work
// to a pool of child goroutines waiting for work to perform.
func pooling() {

	// GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously
	// an argument of 0 returns the current number of logical CPUs used by the Go runtime.

	// NumCPU returns the number of logical CPUs usable by the current process.
	// This is the total number of logical CPUs on the machine, regardless of the current GOMAXPROCS setting.
	// This value can be useful when you're trying to make decisions about the level of concurrency in your program
	// or understand the hardware capabilities of the machine where the process is running.

	g := runtime.GOMAXPROCS(0)
	fmt.Println("num CPU's:", runtime.NumCPU())

	ch := make(chan int)

	for i := 0; i < g; i++ {
		go func(worker int) {
			for d := range ch {
				fmt.Printf("recv: %d - worker %d \n", d, worker)
			}
			fmt.Printf("shutdown signal - worker %d \n", worker)
		}(i)
	}

	// send out work
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("sent: %d \n", i)
	}

	close(ch)
	fmt.Println("sent shutdown signal")

	time.Sleep(2 * time.Second)
	fmt.Println("-------------------------------------------------")
}
