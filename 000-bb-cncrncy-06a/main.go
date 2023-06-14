package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	ii := []int{8, 5, 2, 4, 3, 1, 6, 7, 9}
	fmt.Println(ii)
	ii2 := sleepSort(ii)
	fmt.Println(ii2)
}

func sleepSort(xi []int) []int {
	xi2 := make([]int, 0, len(xi))
	ch := make(chan int)
	sort.Ints(xi)
	go func() {
		for _, v := range xi {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			ch <- v
		}
		close(ch)
	}()

	for v := range ch {
		xi2 = append(xi2, v)
	}
	return xi2
}
