package main

import "fmt"

type human struct {
	name  string
	email string
}

func (u *human) identify() {
	fmt.Printf("%s \t %s \n", u.name, u.email)
}

type secretAgent struct {
	person human // NOT Embedding
	id     string
}

func main() {
	ad := secretAgent{
		person: human{
			name:  "James Bond",
			email: "jbond@uk.gov",
		},
		id: "007",
	}
	ad.person.identify()
}
