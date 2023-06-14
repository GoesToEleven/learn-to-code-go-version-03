package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- 2
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	// cancel is a function: it allows you to manually cancel a context

	select {
	case v := <-ch1:
		fmt.Println("ch1:", v)
	case v := <-ch2:
		fmt.Println("ch2:", v)
	case <-ctx.Done():
		fmt.Println("timeout")
	}

	select {} // blocks forever without consuming CPU
	// for{} loop consumes CPU
}

/*
In Go programming language, `context` is a package that defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.

The context package in Go serves several important functions:

1. Cancellation: You can cancel a context, which will then cancel all the functions that received this context or derived contexts. It's a way to cancel long-running functions or groups of functions.

2. Deadlines and Timeouts: You can set a deadline or a timeout on a context, and all functions using this context will be notified when the time is up.

3. Context Values: You can store key-value pairs in a context. However, this feature should be used with care, as it's not type-safe and can cause difficult-to-diagnose bugs.

The `context.Context` object is often passed as the first argument to a function, and it is generally used in functions that call other functions, which may be running concurrently, or need to be stopped or paused.

Here's an example of how a context could be used in Go:

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    select {
    case <-time.After(2 * time.Second):
        fmt.Println("overslept")
    case <-ctx.Done():
        fmt.Println(ctx.Err())
    }
}
```

In this code, a context with a timeout of 1 second is created. The select statement then waits for either 2 seconds to pass, or for the context to be done (which will happen after 1 second), and then prints a message accordingly. As the context's timeout is shorter, the output will be "context deadline exceeded", indicating that the context ended before `time.After`.

Remember to always call the `cancel` function, or the resources associated with the Context will be leaked. In this example, `cancel` is called with the `defer` statement, ensuring that it will be called when the `main` function returns.

----------

When we say a context will be "leaked," we are referring to a common type of resource leak that occurs in concurrent programming.

In Go, when you create a context with a cancellation function, such as `context.WithCancel()` or `context.WithTimeout()`, an internal system is set up to track and propagate cancellation signals. If you don't call the cancellation function, this system will continue to exist, holding onto resources like memory and potentially preventing garbage collection. Over time, these leaked resources can build up and cause problems such as increased memory usage, decreased performance, or even application crashes.

In other words, not calling the `cancel` function can cause a "context leak," because the resources associated with the context (including memory and any other resources that might have been acquired) will not be released until the program ends or the garbage collector decides to reclaim them.

Here's an example:

```go
ctx, cancel := context.WithCancel(context.Background())

// Some operation that can be cancelled.
go doSomething(ctx)

// If we forget to call this, resources may be leaked.
// cancel()
```

If the `cancel()` function is never called, then the resources associated with `ctx` won't be cleaned up properly, even after `doSomething` has finished executing. That's why it's considered good practice to always defer the `cancel()` call right after creating the context, to ensure the cleanup happens:

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// Some operation that can be cancelled.
go doSomething(ctx)
```

This way, regardless of how the function in which the context was created exits, the `cancel` function will be called, cleaning up the resources associated with the context.
*/
