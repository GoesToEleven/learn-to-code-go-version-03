package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

func doubleInt(n int) int {
	return n*2
}

/*
got test -v

go test -cover

go test -coverprofile c.out

go tool cover -html=c.out

*/