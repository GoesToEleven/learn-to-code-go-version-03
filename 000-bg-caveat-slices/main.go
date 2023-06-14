package main

import "fmt"

func main() {
	xab := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	fmt.Printf("%T\t%v", xab, xab)

	xb := make([][]byte, 0, 10)
	fmt.Println(xb)
	fmt.Println(len(xb))
	fmt.Println(cap(xb))

	// problematic
	// makes a copy of xab and assigns it to v

	// for _, v := range xab {
	// 	fmt.Printf("%T\t%v\n", v[:], v[:])
	// 	xb = append(xb, v[:])
	// 	fmt.Println("appended",xb)
	// }

	// this is good code
	// doesn't make a copy of xab
	// fmt.Println(xab[1])
	
	for i := 0; i < len(xab); i++ {
		n := xab[i]
		xb = append(xb, n[:])		
	}
	fmt.Println("good code",xb)
}
