package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {

	// Cancellation Pattern
	cancellation()
}

// cancellation: In this pattern, the manager goroutine creates a worker
// goroutine to perform some work. The manager goroutine is only willing to
// wait 150 milliseconds for that work to be completed. After 150 milliseconds
// the manager goroutine walks away.
func cancellation() {
	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// IMPORTANT
	// anytime we're doing cancellation with context
	// we need a buffered channel of 1
	// otherwise we can get a goroutine leak
	// if we exceed the cancellation timeout with an *unbuffered* channel
	// the goroutine will block from sending onto the channel
	// and we will have a goroutine leak
	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(r.Intn(200)) * time.Millisecond)
		ch <- "data"
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
