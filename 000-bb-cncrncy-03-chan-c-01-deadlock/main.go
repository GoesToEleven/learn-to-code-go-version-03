package main

import (
	"fmt"
)

// this code DEADLOCKS
// when ranging over a slice or map,
// -- go knows when the range is done
// -- go knows the length of a slice or map
// when ranging over a channel
// -- go does NOT know when items are done being sent
// -- solve this by closing the channel

// send & receive block on unbuffered channels
// until send/receive synchronization occurs

func main() {
	ch := make(chan string)

	go func() {
		for i := 1; i < 4; i++ {
			msg := fmt.Sprintf("message #%d", i)
			ch <- msg
		}
	}()

	for v := range ch {
		fmt.Println("got:", v)
	}

}

// you can do two operations on a channel:
// send
// receive

// CHANNEL SEMANTICS
// send & receive will block until opposite operation (*except buffered channels)
// you can range over a channel and the range will exit when the channel is closed