package main

import "fmt"

type user struct {
	name    string
	surname string
}

func main() {

	users := make(map[string]user)
	users["Mouse"] = user{"Jerry", "Mouse"}
	fmt.Printf("%+v\n", users["Cat"])

	if v, ok := users["Mouse"]; ok {
		fmt.Printf("%+v\n", v)
	}
}
