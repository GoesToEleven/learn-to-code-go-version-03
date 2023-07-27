package main

import "fmt"

func main() {
	qty := 10
	var vs []string
	ps := make([]string, qty)

	vs = valSemantics(vs, qty)
	fmt.Println(len(vs), cap(vs))
	fmt.Printf("%#v \n", vs)

	ptrSemantics(ps, qty)
	fmt.Println(len(ps), cap(ps))
	fmt.Printf("%#v \n", ps)
}

func valSemantics(ss []string, n int) []string {
	for i := 0; i < n; i++ {
		s := fmt.Sprintf("item %d", i)
		ss = append(ss, s)
	}
	return ss
}

func ptrSemantics(ss []string, n int) {
	for i := 0; i < n; i++ {
		ss[i] = fmt.Sprintf("item %d", i)
	}
}

// go test -bench .
