package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func massAdd() int {
	var count int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			count += i + j
		}
	}
	return count
}

func main() {
	fmt.Println(sum(42, 43))
	fmt.Println(multiply(2,3))
	fmt.Println(massAdd())
}

/*
read about "go test"
at the command line, enter "go help"
read the results
then enter "go help test"

to test these files,
at the command line enter any of the following

go test
go test -v
go test .
go test ./...

*/
