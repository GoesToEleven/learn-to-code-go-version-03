package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// sleep takes type duration
	// type duration's underlying type is int64
	d := time.Duration(r.Int63n(250))
	fmt.Printf("%v \t %T\n", d, d)
	time.Sleep(d * time.Millisecond)
	fmt.Println("Done")
}
