package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type customers int

func main() {
	fmt.Println(sumInt(42, 43))
	fmt.Println(sumFloat64(42.03, 43.04))
	fmt.Println(sum[int](42, 43))
	fmt.Println(sum[float64](42.03, 43.04))

	var yesterday customers = 84
	var today customers = 168
	fmt.Println(sum(yesterday, today))
}

func sumInt(x, y int) int {
	return x + y
}
func sumFloat64(x, y float64) float64 {
	return x + y
}

type numTypes interface {
	constraints.Integer | constraints.Float
}

// https://pkg.go.dev/golang.org/x/exp/constraints

func sum[T numTypes](x, y T) T {
	return x + y
}
