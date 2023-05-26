package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	jm := []string{"Miss", "Moneypenny", "I'm 008"}

	xp := [][]string{jb, jm}

	for i, v := range xp {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}
}

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."

Range over the records, then range over the data in each record.

*/
