package main

import "fmt"

func main() {
	fmt.Println("Behold, the full deck of cards!")
	cards := orderedDeck()
	for _, c := range cards {
		fmt.Print(c.String() + " ")
	}
	fmt.Printf("\nNumber of cards: %d\n", len(cards))
}
