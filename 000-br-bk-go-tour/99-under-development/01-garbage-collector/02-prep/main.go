package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	numCPU := runtime.NumCPU()
	fmt.Printf("Number of CPUs: %d\n", numCPU)

	for i := 0; i < 100; i++ {
		s := shareUp()
		fmt.Print(*s)
	}
}

//go:noinline
func shareUp() *string {
	t := time.Now()
	s := fmt.Sprintf("%v \t \t goroutines: %v \t cpus: %v \n", t.Format("03:04:05.0000000"), runtime.NumGoroutine(), runtime.NumCPU())
	time.Sleep(10 * time.Millisecond)
	return &s
}
