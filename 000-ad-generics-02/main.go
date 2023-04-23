package main

import "fmt"

func main() {
	fmt.Println(sumInt(42, 43))
	fmt.Println(sumFloat64(42.03, 43.04))
	fmt.Println(sum[int](42, 43))
	fmt.Println(sum[float64](42.03, 43.04))
}

func sumInt(x, y int) int {
	return x + y
}
func sumFloat64(x, y float64) float64 {
	return x + y
}

func sum[T int | float64](x, y T) T {
	return x + y
}
