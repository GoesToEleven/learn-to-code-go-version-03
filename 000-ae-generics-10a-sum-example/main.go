package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4}
	xf := []float64{22.3, 37.7, 55.5}

	fmt.Println(sumInt(xi))
	fmt.Println(sumNum[int](xi))
	fmt.Println(sumFloat64(xf))
	fmt.Println(sumNum(xf)) // type inference
}

func sumInt(ii []int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

func sumFloat64(ff []float64) float64 {
	total := 0.0
	for _, v := range ff {
		total += v
	}
	return total
}

// generics

type number interface {
	int | float64
}

func sumNum[T number](tt []T) T {
	var total T
	for _, v := range tt {
		total += v
	}
	return total
}
