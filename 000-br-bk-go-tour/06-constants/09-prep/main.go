package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

const a = 42
const b = "hello"

const (
	c = "A"
	d = "🚀"
	e = true
	f = 1 << 3
	g = 42.234e9
	h = math.Pi * math.Pi
)

func main() {
	z := 1 << 62
	fmt.Println(addCommas(strconv.FormatUint(uint64(z), 10)))
	fmt.Println(addCommas(strconv.FormatUint(uint64(18446744073709551615), 10)))
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Printf("%.50f \n", h)
	fmt.Printf("%.100f \n", math.Pi)
	fmt.Printf("%T \n", h)

	// Using math/big
	bi := big.NewInt(0)
	fmt.Println("BI", bi)
	bi.SetBit(bi, 255, 1)  // Equivalent to 1 << 255
	fmt.Println("BI", bi)
}

func addCommas(str string) string {
	n := len(str)
	if n <= 3 {
		return str
	}
	var result []byte
	for i, c := range str {
		result = append(result, byte(c))
		remaining := n - i - 1
		if remaining > 0 && remaining%3 == 0 {
			result = append(result, ',')
		}
	}
	return string(result)
}
