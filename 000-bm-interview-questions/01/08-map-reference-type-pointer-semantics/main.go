package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Lennon"] = 40
	m["McCartney"] = 81
	m["Harrison"] = 58
	m["Starr"] = 83
	m["Best"] = 81

	birthday(m, "Best")
	fmt.Println(m["Best"])
}

// pointer semantics
// a map is a reference data type
func birthday(m map[string]int, person string) {
	m[person] += 1
}
