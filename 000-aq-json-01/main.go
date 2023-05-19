package main

import (
	"fmt"
	"log"
	"net/http"
)

// curl -i https://api.github.com/users/GoesToEleven

func main() {
	resp, err := http.Get("https://api.github.com/users/GoesToEleven")
	if err != nil {
		log.Fatalf("error getting request: %s", err)
	}
	defer resp.Body.Close()
	
	ct := resp.Header.Get("Content-Type")
	fmt.Println(ct)

	fmt.Println(resp.Header.Get("server"))
	fmt.Println(resp.Header.Get("date"))
	fmt.Printf("%#v", resp.Header.Get("content-length"))
}