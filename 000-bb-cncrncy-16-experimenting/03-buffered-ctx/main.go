package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("num goroutines - check 1:", runtime.NumGoroutine())

	ctx, cancel := context.WithTimeout(context.Background(), (80 * time.Millisecond))
	defer cancel()

	// ch := make(chan int) // goroutine leak
	ch := make(chan int, 1) // Buffered channel with buffer of one

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch <- 1
		ch <- 2
		// ch <- 3
		fmt.Println("Goroutine has finished sending")
	}()
	fmt.Println("num goroutines - check 2:", runtime.NumGoroutine())

	select {
	case v := <-ch:
		fmt.Println("rcvd", v)
	case <-ctx.Done():
		fmt.Println("context done")
	}
	fmt.Println("num goroutines - check 3:", runtime.NumGoroutine())

	close(ch)
}

/*
In Go, goroutine leaks can happen when you start a goroutine that sends on a channel but there's no corresponding receiver to receive from that channel. This goroutine is now blocked forever, waiting for a receiver to receive its value, effectively leaking. This is a kind of deadlock situation. Buffered channels can help prevent such a situation.

A buffered channel is a channel with a queue. The capacity of the queue is defined during the creation of the channel. When you send to a buffered channel, the send is placed in the queue. A send operation on a buffered channel is only blocked when the queue is full.

By using a buffered channel with a buffer of one, you're allowing a single send to happen without requiring a corresponding receive to happen immediately. The send operation places the value into the buffer and then proceeds without blocking, even if there's no receiver yet. When a receiver does eventually come along, it will take the value from the buffer.

So by using a buffered channel with a buffer of one, you're allowing a goroutine to perform its send operation and then complete its work without blocking, even if there's no receiver ready to receive yet. This helps prevent the goroutine from leaking.

In this example, the goroutine is able to send the value to the channel and then complete its work, even though there's no receiver ready to receive the value immediately. Without the buffer, the send operation would block and the goroutine would leak if no receiver ever came along.

*/
