package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "surf"
	m[2] = "ski"
	m[3] = "swim"
	m[4] = "sail"
	m[5] = "snowboard"
	m[6] = "sumo wrestling"
	for k, v := range m {
		m[k] = "CHANGED"
		fmt.Println(k, v)
	}
	fmt.Println(m)
}
