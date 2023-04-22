package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%b - %d - %v", 1<<20, 1<<20, math.Pow(2, 20))
}
