package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(time.Now())
	json.NewEncoder(os.Stdout).Encode(time.Now())
}
