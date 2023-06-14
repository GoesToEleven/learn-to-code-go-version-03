package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)

	for v := range c {
		fmt.Println(v)
	}
}

func merge(a, b <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		adone, bdone := false, false
		for !adone || !bdone {
			select {
			case v, ok := <-a:
				if !ok {
					adone = true
					log.Printf("a is done")
					continue
				}
				ch <- v
			case v, ok := <-b:
				if !ok {
					bdone = true
					log.Printf("b is done")
					continue
				}
				ch <- v
			}
		}
	}()
	return ch
}

func asChan(xi ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range xi {
			ch <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(ch)
	}()
	return ch
}

// CHANNEL SEMANTICS
// send & receive will block until opposite operation (*except buffered channels)
// you can range over a channel and the range will exit when the channel is closed
// receive from a closed channel will return the zero value without blocking
// you can check whether a channel is open with the 'comma ok' idiom
// send to a closed channel will panic
// closing a closed channel will panic
// send/receive to a nil channel will block forever

// https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308

/*
In Go, a channel is a conduit that you can use to send and receive values with the channel operator `<-`. Channels are created with the `make` function. The zero value for a channel type is `nil`.

A nil channel in Go is essentially a channel that doesn't yet reference a make-created channel in memory. You can declare a nil channel as follows:

```go
var ch chan int
```

Here, `ch` is a nil channel because it hasn't been initialized with `make`.

When it comes to behavior, a nil channel is always blocked. This means that if you try to send data to a nil channel, the send operation will block forever:

```go
ch <- 1  // This will block forever because ch is nil
```

And if you try to receive data from a nil channel, the receive operation will also block forever:

```go
value := <-ch  // This will block forever because ch is nil
```

It's important to note that using a nil channel can lead to deadlocks and should be used with care.

Now, why might you want to use a nil channel?

Nil channels can be useful in scenarios where you want to disable a send or receive operation in a `select` statement. The `select` statement in Go can be used to choose from multiple send/receive channel operations. If one of these operations is not needed, you can set the corresponding channel to nil, which effectively disables that case in the `select` statement because sends and receives from a nil channel block indefinitely. Here's an example:

```go
var sendCh chan int  // This is a nil channel

select {
case sendCh <- 1:
    fmt.Println("Sent value to sendCh")
case value := <-recvCh:
    fmt.Println("Received value from recvCh:", value)
}
```

In this `select` statement, the case `sendCh <- 1` is effectively disabled because `sendCh` is a nil channel. Only the case reading from `recvCh` (assuming it's a valid, initialized channel) can proceed. This way, you can control the available send/receive operations dynamically.

Just remember that nil channels are not a beginner concept and they require a good understanding of channels and goroutines to avoid potential deadlock situations. Always remember to properly initialize and close your channels where necessary, and be aware that nil channels can always block if you try to send or receive from them.
*/
