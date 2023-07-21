package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[string]int)
	m["Jagger"] = 79
	m["Richards"] = 79
	m["Woods"] = 76
	m["Jones"] = 27

	printM(m)

	// output sorted results
	fmt.Println("***********")
	printSorted(m)
}

func printM(m map[string]int) {
	for i := 0; i < 20; i++ {
		var s string
		for k := range m {
			s += fmt.Sprintf("%v-%d ", k, m[k])
		}
		fmt.Println(s)
	}
}

func printSorted(m map[string]int) {
	xs := make([]string, 0, len(m))
	for k := range m {
		xs = append(xs, k)
	}
	sort.Strings(xs)

	for i := 0; i < 20; i++ {
		var s string
		for _, v := range xs {
			s += fmt.Sprintf("%v-%d ", v, m[v])
		}
		fmt.Println(s)
	}
}
