package main

import "fmt"

func main() {

	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Adventure"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"dr-no":           {"Needing love", "Ice Cream", "Sunsets"},
	}

	for i, xs := range m {
		fmt.Println("name: ", i)
		for j, val := range xs {
			fmt.Printf("\t index: %v \t value: %v\n", j, val)
		}
	}
}

// source: student, Sergey
