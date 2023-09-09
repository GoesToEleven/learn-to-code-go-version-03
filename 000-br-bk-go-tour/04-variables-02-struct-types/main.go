package main

import (
	"fmt"
	"unsafe"
)

func main() {
	example1()
	example2()
}

func example1() {
	type Example struct {
		A bool    // 1 byte
		B int32   // 4 bytes
		C float64 // 8 bytes
	}

	fmt.Println("EXAMPLE 1")
	fmt.Println("Alignment of bool:", unsafe.Alignof(bool(true)))
	fmt.Println("Alignment of int32:", unsafe.Alignof(int32(0)))
	fmt.Println("Alignment of float64:", unsafe.Alignof(float64(0.0)))
	fmt.Println("Alignment of Example:", unsafe.Alignof(Example{}))
	fmt.Println()
}

func example2() {
	// Without optimization
	type Unoptimized struct {
		A bool    // 1 byte
		B float64 // 8 bytes
		C int32   // 4 bytes
		// Padding: 3 bytes for A, 4 bytes for C (assuming 64-bit machine)
	}
	b := Unoptimized{}
	fmt.Println("EXAMPLE 2 - Unoptimized")
	fmt.Println("Alignment of bool:", unsafe.Alignof(bool(true)))
	fmt.Println("Alignment of float64:", unsafe.Alignof(float64(0.0)))
	fmt.Println("Alignment of int32:", unsafe.Alignof(int32(0)))
	fmt.Println("Alignment of Unoptimized:", unsafe.Alignof(b))
	fmt.Println("Size of Unoptimized:", unsafe.Sizeof(b))
	fmt.Println()

	// With optimization
	type Optimized struct {
		B float64 // 8 bytes
		C int32   // 4 bytes
		A bool    // 1 byte
		// Padding: None or minimal (assuming 64-bit machine)
	}
	c := Optimized{}
	fmt.Println("EXAMPLE 2 - Optimized")
	fmt.Println("Alignment of float64:", unsafe.Alignof(float64(0.0)))
	fmt.Println("Alignment of int32:", unsafe.Alignof(int32(0)))
	fmt.Println("Alignment of bool:", unsafe.Alignof(bool(true)))
	fmt.Println("Alignment of Optimized:", unsafe.Alignof(c))
	fmt.Println("Size of Optimized:", unsafe.Sizeof(c))
	fmt.Println()
}
