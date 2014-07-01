package main

import "fmt"

type runner struct {
	state   state
	players []player
	hands   []hand
}

type state struct {
	history         history
	scores          []int
	heartsAreBroken bool
}

type history []trick

type trick []move

type move struct {
	player int
	card   card
}

type hand []card

type player interface {
	swapCards(hand) []int
	move(state, hand) int
}

func newRunner(p0, p1, p2, p3 player) {
	r := runner{}
	r.players = make([]player, 4)
	r.players[0] = p0
	r.players[1] = p1
	r.players[2] = p2
	r.players[3] = p3

	d := shuffledDeck()
	r.hands = make([]hand, 4)
	r.hands[0] = d[0:13]
	r.hands[1] = d[13:26]
	r.hands[2] = d[26:39]
	r.hands[3] = d[39:52]

	for _, h := range r.hands {
		for _, c := range h {

			fmt.Print(c.String(), " ")
		}
		fmt.Printf("  %d", len(h))
		fmt.Println("")
	}
}
