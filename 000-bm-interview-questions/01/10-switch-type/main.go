package main

import "fmt"

func main() {

	// fmt.Println can be called with values of any type.
	fmt.Println("Hello, world")
	fmt.Println(12345)
	fmt.Println(3.14159)
	fmt.Println(true)

	// How can we do the same?
	myPrintln("Hello, world")
	myPrintln(12345)
	myPrintln(3.14159)
	myPrintln(true)

	var s any = "costa mesa"
	if v, ok := s.(string); ok {
		fmt.Printf("%T \t %v \n", v, v)
		fmt.Println(ok)
	}

}

func myPrintln(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Printf("Is string  : type(%T) : value(%s)\n", v, v)
	case int:
		fmt.Printf("Is int     : type(%T) : value(%d)\n", v, v)
	case float64:
		fmt.Printf("Is float64 : type(%T) : value(%f)\n", v, v)
	default:
		fmt.Printf("Is unknown : type(%T) : value(%v)\n", v, v)
	}
}
