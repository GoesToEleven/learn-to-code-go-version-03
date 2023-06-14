package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")

	ch := make(chan int)
	sig := make(chan int)

	go producer(ch, sig)

	go puller(ch, 1)

	go puller(ch, 2)

	fmt.Println("from sig", <-sig)
	time.Sleep(2 * time.Second)
	close(ch)
	close(sig)
}

func producer(ch chan<- int, sig chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	sig <- 1
}

func puller(ch <-chan int, signal int) {
	for v := range ch { // when there's a value to pull, it pulls it
		fmt.Println("from:", signal, v)
	}
}
