package main

import (
	"fmt"
	"mymodule/000-bm-interview-questions/01/14-packaging/01/users"
	"mymodule/000-bm-interview-questions/01/14-packaging/01/users/admins"
)

func main() {
	fmt.Println(users.Hello())
	fmt.Println(admins.Hello())
}
