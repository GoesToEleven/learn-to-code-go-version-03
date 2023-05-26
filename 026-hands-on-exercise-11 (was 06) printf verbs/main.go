package main

import "fmt"

func main() {
	s, i, f := "Happiness", 42, 42.42
	fmt.Printf("%v is of type %T\n",i,i)
	fmt.Printf("%v is of type %T\n",f,f)
	fmt.Printf("%v is of type %T\n",s,s)
}