package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m["tunde"])
	m = make(map[string]int)
	m["todd"] = 42
	m["henry"] = 16

	//  --- or ---

	// m := make(map[string]int)
	// m["todd"] = 42
	// m["henry"] = 16

	//  --- or ---

	//	m := map[string]int{
	// 	"todd": 42,
	// 	"henry": 16,
	// }

	for k := range m {
		fmt.Printf("just the keys: %s\n", k)
	}

	for k, v := range m {
		fmt.Printf("%s is %d years old\n", k, v)
	}

	for _, v := range m {
		fmt.Printf("just the values: %d\n", v)
	}

	if v, ok := m["Padget"]; ok {
		fmt.Printf("%s is %d years old\n", "Padget", v)
	} else {
		fmt.Printf("%s not found\n", "Padget")
	}

	//delete
	m["Shakespeare"] = 459
	fmt.Printf("%#v\n", m)
	delete(m, "Shakespeare")
	fmt.Printf("%#v\n", m)
	delete(m, "Shakespeare") // no panic

	// len
	fmt.Println("len: ", len(m))
}
