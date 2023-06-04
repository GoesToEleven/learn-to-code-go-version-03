package main

import "fmt"

func main() {
	fmt.Println(Paradise("Hawaii"))
}

// Paradise prints out the user's input of paradise as a sentence.
func Paradise(loc string) string {
	return fmt.Sprint("My idea of paradise is ", loc)
}
