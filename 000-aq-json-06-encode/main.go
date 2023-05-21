package main

import (
	"encoding/json"
	"os"
	"time"
)

func main() {
	json.NewEncoder(os.Stdout).Encode(time.Now())
}
