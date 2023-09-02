package main

import (
	"fmt"
	"mymodule/000-bm-interview-questions/01/14-packaging/01-hierarchy-no-sub-packaging/users"
	"mymodule/000-bm-interview-questions/01/14-packaging/01-hierarchy-no-sub-packaging/users/admins"
)

func main() {
	fmt.Println(users.Hello())
	fmt.Println(admins.Hello())
}
