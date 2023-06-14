package main

import (
	"fmt"
)

// this code DOES run

// send & receive block on unbuffered channels 
// until send/receive synchronization occurs

func main() {
	ch := make(chan string)
	go func() {
		// sending on an unbuffered channel blocks
		ch <- "hi"
	}()

	// receiving an unbuffered channel blocks
	msg := <-ch

	fmt.Println(msg)
}

// you can do two operations on a channel:
// send
// receive

// CHANNEL SEMANTICS
// send & receive will block until opposite operation (*except buffered channels)