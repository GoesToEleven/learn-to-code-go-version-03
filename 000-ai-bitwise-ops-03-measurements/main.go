package main

import "fmt"

type ByteSize int

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	// ZB
	// YB
)

func main() {
	fmt.Printf("KB %d \t\t %b\n", KB, KB)
	fmt.Printf("MB %d \t\t %b\n", MB, MB)
	fmt.Printf("GB %d \t\t %b\n", GB, GB)
	fmt.Printf("TB %d \t %b\n", TB, TB)
	fmt.Printf("PB %d \t %b\n", PB, PB)
	fmt.Printf("EB %d \t %b\n", EB, EB)
}

/*
PB		    1125899906842624
EB		 1152921504606846976 
int		18446744073709551615
*/