package main

import "fmt"

func main() {
	s := []string{"surf", "skis", "snow", "swim"}
	for i, v := range s {
		fmt.Printf("%v - %p \t %v - %p \n", s[i], &s[i], v, &v)
	}
}
