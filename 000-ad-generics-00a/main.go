package main

import "fmt"

func main() {
	loveIt("Hello world")
	loveIt(7)
	loveIt(true)
	loveIt(4.2)
}

func loveIt[T any](a T) {
	fmt.Println(a)
}
