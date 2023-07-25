package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	waitForResult()
}

func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(r.Intn(500)) * time.Millisecond)
		ch <- "data"
		fmt.Println("sent signal")
	}()

	d := <-ch
	fmt.Println("recv'd signal :", d)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

/*
Unbuffered channels don't have any capacity to hold data,
so they require both a sender and receiver to be ready at the same time.
The send operation on an unbuffered channel blocks until
another goroutine is ready to receive the data,
and conversely, the receive operation blocks
until another goroutine is ready to send data.
*/
