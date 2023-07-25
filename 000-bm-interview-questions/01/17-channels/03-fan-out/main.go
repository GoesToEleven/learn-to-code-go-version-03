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
	fanOut()
}

func fanOut() {
	n := 10
	ch := make(chan string)

	for c := 0; c < n; c++ {
		go func(child int) {
			time.Sleep(time.Duration(r.Intn(200)) * time.Millisecond)
			ch <- "data"
			fmt.Println("child : sent signal :", child)
		}(c)
	}

	for n > 0 {
		d := <-ch
		n--
		fmt.Println(d)
		fmt.Println("parent : recv'd signal :", n)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
