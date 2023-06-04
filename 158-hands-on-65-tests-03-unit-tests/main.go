package main

import "fmt"

func main() {
	fmt.Println(paradise("Hawaii"))
}

// paradise returns a string
func paradise(loc string) string {
	return fmt.Sprint("My idea of paradise is ", loc)
}
