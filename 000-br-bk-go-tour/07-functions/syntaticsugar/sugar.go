package syntaticsugar

import (
	"fmt"
	"log"
)

type user struct {
	first string
}

// ffUser is only returning a pointer for teaching purposes: to illustrate syntactic sugar.
// In production code, there is no need to return a pointer here.
func ffUser(s string) (*user, error) {
	u := user{s}
	return &u, nil
}

func RunSugar() {

	u, err := ffUser("Jenny")
	if err != nil {
		log.Fatalf("error creating user %s \n", err)
	}

	// address of a struct
	fmt.Println(u)
	fmt.Printf("%T \n", u)
	fmt.Println((*u).first)

	// syntatic sugar
	fmt.Println(u.first)
}
