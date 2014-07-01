package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Behold, the full deck of cards!")
	cards := shuffledDeck()
	for _, c := range cards {
		fmt.Print(c.String() + " ")
	}
	fmt.Printf("\nNumber of cards: %d\n", len(cards))

	fmt.Println("")
	newRunner(nil, nil, nil, nil)
}
