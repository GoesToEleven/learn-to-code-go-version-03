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
		close(ch)
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


// YOU DON'T HAVE TO CLOSE A CHANNEL
/*
**Why Do You Close a Channel?**

Closing a channel in Go is used to signal that no more values will be sent on it. This is useful when you want to communicate to the receiver that there are no more values coming, which can be important in the case of range loops, or in general when a goroutine is consuming values and needs to know when to stop.

Here's an example:
```go
ch := make(chan int)
go func() {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)
}()

for i := range ch {
    fmt.Println(i)  // Prints numbers 0 through 9.
}
```
In the example above, once the channel is closed, the `range` loop ends.

**Do You Have to Close a Channel?**

It's not mandatory to close channels. The garbage collector will reclaim the resources used by a channel when it's no longer accessible, regardless of whether it's closed or not. 

You usually want to close a channel to let the receiving goroutine know that no more data will be sent on the channel. If you don't close the channel and the receiver is using a `for... range` loop on the channel, then the loop will never terminate.

However, in a scenario where the sender is a goroutine that you know will eventually end, and the receiver is using a select statement with a default case, it may not be necessary to close the channel.

**What Happens to the Computer's Resources If You Don't Close a Channel and the Program Ends?**

When a Go program terminates, all resources (including memory, file descriptors, channels etc.) associated with the program are cleaned up by the operating system. This is not specific to Go, but is a standard behavior of modern operating systems.

Therefore, if a Go program ends normally or due to an unhandled exception, any open channels or other resources will be released. It's still a good practice to close resources when they're no longer needed, but if your program ends (either because it completed its work or because an error occurred) any resources it was using will be cleaned up.

*/