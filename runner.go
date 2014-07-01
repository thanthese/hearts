package main

// "round" is a collection of tricks
// "deal"

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
	swapCards
}
