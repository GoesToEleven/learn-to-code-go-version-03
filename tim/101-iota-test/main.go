package main

import "fmt"

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println("BunchE")
	type Bunch[E any] []E
	b := Bunch[int]{1, 2, 3}
	fmt.Printf("b is of type %#v and with value %v \n", b, b)
	fmt.Println("MB")
	fmt.Printf("b is of type %#v and with value %v \n", MB, MB)
	fmt.Println("KB")
	fmt.Printf("b is of type %#v and with value %v \n", KB, KB)
}
