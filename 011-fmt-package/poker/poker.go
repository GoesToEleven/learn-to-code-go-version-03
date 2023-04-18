package poker

import (
	"fmt"
	"math/rand"
)

func Deal() {
	face := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	suite := []string{"❤️", "♦", "♠", "♣"}

	for i := 0; i < 5; i++ {
		fmt.Println(face[rand.Intn(len(face))], suite[rand.Intn(len(suite))])
	}
}
