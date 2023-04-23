package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Println(sum(42, 43))
	fmt.Println(multiply(2,3))
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
