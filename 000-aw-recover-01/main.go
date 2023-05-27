package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(safeDiv(10, 0))
	fmt.Println(safeDiv(10, 2))
}

func safeDiv(a, b int) (v int, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR:", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return a / b, nil
}
