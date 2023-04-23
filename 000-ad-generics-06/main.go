package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type numTypes interface {
	constraints.Integer | constraints.Float
}

func sumNums[T numTypes](xt ...T) T {
	var result T
	for _, v := range xt {
		result += v
	}
	return result
}

func main() {
	fmt.Println(sumNums(42, 43, 44))
}

// https://pkg.go.dev/golang.org/x/exp/constraints
