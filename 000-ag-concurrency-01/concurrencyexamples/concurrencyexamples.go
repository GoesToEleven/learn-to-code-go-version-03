package concurrencyexamples

import (
	"fmt"
	"sync/atomic"
)

func One() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	x := <-ch
	fmt.Println(x)
}

func Two() {
	ch := make(chan int)

	for i := 0; i < 1000; i++ {
		go func(n int) {
			ch <- n
		}(i)
	}

	var counter int64
	for i := 0; i < 1000; i++ {
		x := <-ch
		fmt.Println(x)
		atomic.AddInt64(&counter, 1)
	}
	fmt.Println("Counter:", atomic.LoadInt64(&counter)) // access without race
}

func Three() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}

func Four() {
	ch := make(chan int, 1) // buffered channel
	ch <- 42
	x := <-ch
	fmt.Println(x)

	ch2 := make(chan int, 2) // buffered channel
	ch2 <- 43
	ch2 <- 44
	y := <-ch2
	fmt.Println(y)
	z := <-ch2
	fmt.Println(z)

	// this will deadlock
	/*
		ch3 := make(chan int, 2) // buffered channel
		ch3 <- 45
		ch3 <- 46
		ch3 <- 47
		a := <-ch3
		fmt.Println(a)
		b := <-ch3
		fmt.Println(b)
	*/
}

func Five() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		ch2 <- 42
	}()

	select {
	case val := <-ch1:
		fmt.Println("value from ch1", val)
	case val := <-ch2:
		fmt.Println("value from ch2", val)
	}
}
