package main

import (
	"fmt"
	"time"
)

func main() {
	_, month, day := time.Now().Date()
	fmt.Println(month)
	fmt.Println(day)
	if month == time.November && day == 10 {
		fmt.Println("Happy Go day!")
	}
}
