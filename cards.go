package main

func orderedDeck() []card {
	cards := make([]card, 0, 52)
	for _, s := range suits {
		for _, r := range ranks {
			cards = append(cards, card{s, r})
		}
	}
	return cards
}

type suit uint
type rank uint

const (
	heart   suit = iota
	spade        = iota
	diamond      = iota
	club         = iota
)

const (
	two   rank = iota
	three      = iota
	four       = iota
	five       = iota
	six        = iota
	seven      = iota
	eight      = iota
	nine       = iota
	ten        = iota
	jack       = iota
	queen      = iota
	king       = iota
	ace        = iota
)

var suits = []suit{heart, spade, diamond, club}
var ranks = []rank{two, three, four, five, six,
	seven, eight, nine, ten, jack, queen, king, ace}

func (s *suit) String() string {
	switch *s {
	case heart:
		return "♡"
	case spade:
		return "♣"
	case diamond:
		return "♢"
	case club:
		return "♠"
	default:
		return "X" // error
	}
}

func (r *rank) String() string {
	switch *r {
	case two:
		return "2"
	case three:
		return "3"
	case four:
		return "4"
	case five:
		return "5"
	case six:
		return "6"
	case seven:
		return "7"
	case eight:
		return "8"
	case nine:
		return "9"
	case ten:
		return "t"
	case jack:
		return "J"
	case queen:
		return "Q"
	case king:
		return "K"
	case ace:
		return "A"
	default:
		return "X" // error
	}
}

type card struct {
	suit suit
	rank rank
}

func (c *card) String() string {
	return c.rank.String() + c.suit.String()
}
