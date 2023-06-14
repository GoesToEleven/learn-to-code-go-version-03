package main

import (
	"fmt"
	"time"
)

func main() {
	ii := []int{8, 5, 2, 4, 3, 1, 6, 7, 9}
	fmt.Println(ii)

	ii2 := sleepSort(ii)
	fmt.Println(ii2)
}

func sleepSort(xi []int) []int {
	out := make([]int, 0, len(xi))
	ch := make(chan int)

	// FAN OUT
	for _, v := range xi {
		go func(n int) {
			time.Sleep(time.Duration(n*50) * time.Millisecond)
			ch <- n
		}(v)
	}
	

	// this is interesting
	// you can range loop over a slice and not return an index or value
	for range xi {
		out = append(out, <-ch)
	}
	return out
}

// THIS IS NOT GOOD CODE - just interesting