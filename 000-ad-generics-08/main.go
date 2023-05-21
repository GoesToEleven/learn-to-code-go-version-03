package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4}
	xf := []float64{1, 2, 3, 4}

	fmt.Println(maxInts(xi))
	fmt.Println(maxFloat64(xf))
	fmt.Println(max(xi))
	fmt.Println(max(xf))
}

func maxInts(s []int) int {
	if len(s) == 0 {
		return 0
	}
	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func maxFloat64(s []float64) float64 {
	if len(s) == 0 {
		return 0
	}
	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func max[T int | float64](s []T) T {
	if len(s) == 0 {
		return 0
	}
	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// or like this
/*

type Number interface {
	int | float64
}

func max[T Number](s []T) T {
	if len(s) == 0 {
		return 0
	}
	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

*/
