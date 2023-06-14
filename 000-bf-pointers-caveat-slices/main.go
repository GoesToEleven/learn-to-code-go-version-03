package main

import "fmt"

type user struct {
	name  string
	count int
}

func (u *user) addTo() {
	u.count++
}

func main() {
	xu := []user{{"mickey", 0}, {"goofy", 0}}
	fmt.Println(xu)

	mickey := &xu[0] // be careful when using the memory address of a slice
	fmt.Println("pointer to mickey", *mickey)

	xu = append(xu, user{"donald duck", 0})
	mickey.addTo()
	fmt.Println("end result 1", *mickey)
	fmt.Println("end result 2", xu)
}
