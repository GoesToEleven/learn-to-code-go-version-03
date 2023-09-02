package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a time in UTC
	utcTime := time.Date(2023, time.September, 2, 14, 0, 0, 0, time.UTC)
	fmt.Printf("UTC time: \t %v \n", utcTime)

	// Define EST (UTC-5)
	est := time.FixedZone("EST", -5*60*60)

	// Convert the time to EST
	estTime := utcTime.In(est)
	fmt.Printf("Time in EST: \t %v \n", estTime)
}
