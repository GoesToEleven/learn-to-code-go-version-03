package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count atomic.Int64
	fmt.Printf("%T - %v\n", count.Load(), count.Load())

	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10_000; j++ {
				count.Add(1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(count.Load())
}

// https://go.dev/doc/articles/race_detector#Runtime_Overheads

// synchronization vs orchestration

/*
In Go programming language, both synchronization and orchestration are important concepts related to managing the execution flow of concurrent or parallel tasks, with some subtle differences between them.

1. **Synchronization:** 

Synchronization in Go is about making sure multiple goroutines can access shared state without causing race conditions, data inconsistencies, or unexpected behavior. This is commonly achieved using various synchronization primitives provided by the Go standard library, such as `Mutex`, `RWMutex`, and `WaitGroup`.

- `Mutex` or `RWMutex` allows you to lock access to a shared resource to ensure only one goroutine can manipulate the resource at a time, preventing race conditions.
- `WaitGroup` allows you to wait for a collection of goroutines to finish executing before moving on, which is useful when you want to launch multiple goroutines but need to wait until all have completed before proceeding.

For example, if you have multiple goroutines incrementing a shared counter, you could use a `Mutex` to synchronize access to the counter:

```go
var (
    mutex sync.Mutex
    counter int
)

func increment() {
    mutex.Lock()
    counter++
    mutex.Unlock()
}
```

2. **Orchestration:**

Orchestration in Go involves managing and coordinating the execution, scheduling, and communication between goroutines. This typically includes determining the order in which goroutines are run, passing data between them, and handling errors. Channels in Go are a primary tool for achieving this.

- You can use unbuffered channels to synchronize the execution of goroutines or to pass a single piece of data between them.
- Buffered channels allow you to create a pipeline of goroutines where each goroutine can work on a piece of data, then pass it on to the next goroutine in the pipeline.
- Select statements provide a way to handle situations where you have multiple channels and you need to handle whichever one is ready first.

For instance, to coordinate two goroutines to work in a producer-consumer pattern, you can use channels:

```go
func main() {
    ch := make(chan int)
    
    go func() {  // producer
        for i := 0; i < 10; i++ {
            ch <- i
        }
        close(ch)
    }()
    
    go func() {  // consumer
        for n := range ch {
            fmt.Println(n)
        }
    }()
    
    time.Sleep(time.Second)
}
```

In summary, synchronization is about ensuring safe access to shared state, while orchestration is about coordinating the execution of goroutines and passing data between them.

*/