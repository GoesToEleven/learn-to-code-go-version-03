package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(100 * time.Millisecond)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("start", start)
	fmt.Println("now", t)
	fmt.Println("elapsed", elapsed)
}
