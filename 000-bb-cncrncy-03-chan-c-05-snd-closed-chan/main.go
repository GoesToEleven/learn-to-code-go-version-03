package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		for i := 1; i < 4; i++ {
			msg := fmt.Sprintf("message #%d", i)
			ch <- msg
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println("got:", v)
	}

	msg, ok := <-ch
	fmt.Printf("closed: %#v \t ok=%v\n", msg, ok)

	// send on a closed channel will PANIC
	ch <- "Hi"
}

// CHANNEL SEMANTICS
// send & receive will block until opposite operation (*except buffered channels)
// you can range over a channel and the range will exit when the channel is closed
// receive from a closed channel will return the zero value without blocking
// you can check whether a channel is open with the 'comma ok' idiom
// send to a closed channel will panic
