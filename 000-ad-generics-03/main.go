package main

import "fmt"

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

func sum[T ~int | ~float64](x, y T) T {
	return x + y
}
