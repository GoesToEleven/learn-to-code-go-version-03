package main

import (
	"fmt"
)

// this code DOES NOT run

func main() {
	ch := make(chan string)
	ch <- "hi"  // send
	msg := <-ch // receive
	fmt.Println(msg)
}

// you can do two operations on a channel:
// send
// receive

// CHANNEL SEMANTICS
// send & receive will block until opposite operation (*except buffered channels)