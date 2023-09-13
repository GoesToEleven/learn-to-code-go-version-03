package main

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
)

const a = 42
const b = "hello"

const (
	c = "C"
	d = "🚀"
	e = true
	f = 1 << 3
	g = 42.234e9
	h = math.Pi * math.Pi
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%b \n", f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Printf("%T \n", h)
	fmt.Printf("%.50f \n", h)
	fmt.Printf("%.100f \n", math.Pi)

	// Untyped Constants
	const i = 3    // implicit conversion ("kind"): integer
	const j = 3.14 // implicit conversion ("kind"): floating-point
	fmt.Println("i", i)
	fmt.Println("j", j)

	// numerical literals are untyped constants ("constants of a kind")
	// the compiler determines their type at compile time

	// Typed Constants
	// still use the constant type system
	// but their precision is restricted
	const k int = 12345        // type: int
	const l float64 = 3.141592 // type: float64
	fmt.Println("k", k)
	fmt.Println("l", l)

	// this will ERROR: constant 1000 overflows uint8
	// const x uint8 = 1000

	// the compiler does implicit conversaion at compile time
	// implicit because it is not "explicit" - we have not done it

	var m = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)
	fmt.Println(m)

	const n = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)
	fmt.Println(n)

	// const o = 1 / 3 // KindInt(1) / KindInt(3)
	// fmt.Println("--------WE ARE HERE 1 -------")
	// fmt.Println(o)

	// typed and untyped constants must have like types to perform math.
	const y int8 = 1
	const z = 2 * y // int8(2) * int8(1)
	fmt.Println("--------WE ARE HERE-------")
	fmt.Printf("%T \t %v \n", z, z)

	const aaa = 42
	const bbb = 43.01
	fmt.Println("--------KIND-------")
	fmt.Println(reflect.ValueOf(aaa).Kind())
	fmt.Println(reflect.ValueOf(bbb).Kind())
	fmt.Println(reflect.ValueOf(aaa * bbb).Kind())
	fmt.Println("--------TYPE-------")
	fmt.Printf("%T \n", aaa)
	fmt.Printf("%T \n", bbb)
	fmt.Printf("%T \n", aaa*bbb)

	fmt.Println("-------- FLOAT ME -------")
	fm := floatMe(30)
	fmt.Printf("%T \n", fm)

	fmt.Println("-------- Max Int -------")
	// Max integer value on 64 bit architecture
	const maxInt = 9223372036854775807
	fmt.Println("MAXINT", maxInt)
	fmt.Println("MAXINT", addCommas(strconv.Itoa(maxInt)))
	fmt.Printf("MAXINT \t %T \n", maxInt)

	// Much larger value than int64
	const bigger = 9223372036854775808543522345
	fmt.Println("BIGGER", bigger*.1)
	fmt.Printf("BIGGER \t %T \n", bigger*.1)

	// ERROR: typed constant int64 creates overflow
	// const biggerInt int64 = 9223372036854775808543522345

	fmt.Println("-------- iota -------")
	const (
		aa = iota
		bb = iota
		cc = iota
	)

	const (
		dd = iota
		ee
		ff
	)

	fmt.Println("iota ", aa)
	fmt.Println("iota ", bb)
	fmt.Println("iota ", cc)
	fmt.Println("iota ", dd)
	fmt.Println("iota ", ee)
	fmt.Println("iota ", ff)

	fmt.Printf("%b \t %d \n", 1<<3, 1<<3)
	/*
		1 << 10
		This is a bitwise left shift operation.
		The number 1 is represented in binary as 0001.
		Shifting it to the left 10 times will move the 1 ten places to the left,
		filling the empty spaces with zeros. The result in binary will be 10_000_000_000, which in decimal is 1024.
	*/

	const (
		_      = iota // ignore first value by assigning to blank identifier
		KB int = 1 << (10 * iota)
		MB
		GB
		TB
	)
	fmt.Printf("KB \t %d \t\t\t %b \n", KB, KB)
	fmt.Printf("MB \t %d \t\t %b \n", MB, MB)
	fmt.Printf("GB \t %d \t\t %b \n", GB, GB)
	fmt.Printf("TB \t %d \t\t %b \n", TB, TB)

	fmt.Println("-------- math/big -------")
	// Using math/big
	bi := big.NewInt(0)
	fmt.Println("BI", bi)
	bi.SetBit(bi, 255, 1) // Equivalent to 1 << 255
	fmt.Println(bi)
	fmt.Println(addCommas(bi.String()))
}

func floatMe(f float64) float64 {
	return f * 2.0
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
