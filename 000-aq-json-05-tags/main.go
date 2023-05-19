package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// command palette (ctrl+shift+p): Go: add tags to struct fields 

func main() {
	resp, err := http.Get("https://api.github.com/users/GoesToEleven")
	if err != nil {
		log.Fatalf("error getting request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}

	var u User
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&u); err != nil {
		log.Fatalf("error decoding: %s", err)
	}
	fmt.Println("Name\t\t", u.Name)
	fmt.Println("Public Repos\t", u.NumRepos)
	fmt.Println("----")
	fmt.Printf("%#v", u)
}

type User struct {
	Name     string `json:"name,omitempty"`
	NumRepos int    `json:"num_repos,omitempty"`
}
