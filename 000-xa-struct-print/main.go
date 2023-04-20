package main

import "fmt"

func main() {

	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Sciense"},
		"no_dr":           {"Being evil", "Ice Cream", "Sunsets"},
	}

	for i, slice := range m {
		fmt.Println("name: ", i)
		for j, val := range slice {
			fmt.Printf("\t index: %v \t value: %v\n", j, val)
		}
	}
}

// source: student, Sergey
