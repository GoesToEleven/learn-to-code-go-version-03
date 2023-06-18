package main

import (
	"fmt"
)

type ByteSize float64

const (
	_ int = 1 << (10*iota) 	// iota = 0
	KB int = 1 << (10*iota) 	// iota = 1
	MB 
	GB 
	TB 
	PB 
	EB 
	//ZB 
	//YB 
)

func main() {
	fmt.Println("201-hands-on-9\n")
	fmt.Printf("%d \t %b\n", KB, KB)
	fmt.Println("MB")
	fmt.Printf("%d \t %b\n", MB, MB)
	fmt.Printf("%d \t %b\n", GB, GB)
	fmt.Printf("%d \t %b\n", TB, TB)
	fmt.Printf("%d \t %b\n", PB, PB)
	fmt.Printf("%d \t %b\n", EB, EB)
	//fmt.Printf("%d \t %b\n", ZB, ZB)
	//fmt.Printf("%d \t %b\n", YB, YB)
}
