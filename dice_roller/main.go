package main

import (
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func parseInput(input string) (int, int, int, error) {
	diceExpr := regexp.MustCompile(`^(\d+)d(\d+)([+-]\d+)?$`)
	matches := diceExpr.FindStringSubmatch(input)

	if matches == nil {
		return 0, 0, 0, fmt.Errorf("invalid input format")
	}

	numDice, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, 0, 0, err
	}

	numSides, err := strconv.Atoi(matches[2])
	if err != nil {
		return 0, 0, 0, err
	}

	modifier := 0
	if len(matches) > 3 {
		modifier, err = strconv.Atoi(matches[3])
		if err != nil {
			return 0, 0, 0, err
		}
	}

	return numDice, numSides, modifier, nil
}

func rollDice(numDice, numSides int, adv, dis *bool) int {
	total := 0
	for i := 0; i < numDice; i++ {
		roll := rand.Intn(numSides) + 1
		if *adv {
			roll = max(roll, rand.Intn(numSides)+1)
		} else if *dis {
			roll = min(roll, rand.Intn(numSides)+1)
		}
		fmt.Printf("Die %d: %d\n", i+1, roll)
		total += roll
	}
	return total
}

func main() {
	adv := flag.Bool("adv", false, "Roll twice and take the higher")
	dis := flag.Bool("dis", false, "Roll twice and take the lower")
	flag.Parse()

	fmt.Println("Enter dice expression ([number of dice]d[size of dice]+/-[modifier] ie:1d20+4): ")
	var input string
	fmt.Scanln(&input)

	numDice, numSides, modifier, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	total := rollDice(numDice, numSides, adv, dis)
	total += modifier
	fmt.Printf("Total: %d (including modifier of %d)\n", total, modifier)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
