package main

import "fmt"

type anyNum interface {
	int | float64 | string
}

func printAny[T anyNum](values []T){
	for _, v := range values {
		fmt.Print(v, " ")
	}
}


func main(){
	printAny([]int{1,2,3})
	printAny([]string{"a","b","c"})
}

