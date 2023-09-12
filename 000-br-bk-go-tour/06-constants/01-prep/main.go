package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

func main() {
	a := "Hello"
	fmt.Println(reflect.ValueOf(a).Kind())

	var i int
	var f float64
	var c complex128
	var s string
	var p *int

	fmt.Printf("Size of int: %d bytes\n", unsafe.Sizeof(i))
	fmt.Printf("Size of float64: %d bytes\n", unsafe.Sizeof(f))
	fmt.Printf("Size of complex128: %d bytes\n", unsafe.Sizeof(c))
	fmt.Printf("Size of string: %d bytes\n", unsafe.Sizeof(s))
	fmt.Printf("Size of pointer: %d bytes\n", unsafe.Sizeof(p))

	const b = math.Pi
	fmt.Printf("Size of pointer: %d bytes\n", unsafe.Sizeof(b))

	fmt.Printf("%b \t %d \n", 1 << 2, 1 << 2)
	fmt.Printf("%b \t %d \n", 1 << 30, 1 << 30)
}
