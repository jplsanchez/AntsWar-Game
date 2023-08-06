package game

import (
	"math/rand"
	"time"
)

const Cards_Per_Team = 19

type Team uint8

const (
	TeamRed Team = iota
	TeamBlack
)

func (t Team) Enemy() Team {
	if t == TeamBlack {
		return TeamRed
	}
	if t == TeamRed {
		return TeamBlack
	}
	panic("Unknown team")
}

type Card struct {
	Value int
	Team  Team
}

type Deck []Card

func NewDeck(team Team) *Deck {
	queen := Card{Value: 0, Team: team}
	deck := Deck{queen}

	possibleValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, value := range possibleValues {
		deck = append(deck, Card{Value: value, Team: team}, Card{Value: value, Team: team})
	}

	return &deck
}

func (d *Deck) Shuffle(seedOffset ...int64) *Deck {
	seed := time.Now().UnixNano()
	for _, offset := range seedOffset {
		seed += offset
	}
	random := rand.New(rand.NewSource(seed))
	random.Shuffle(len(*d), func(i, j int) { (*d)[i], (*d)[j] = (*d)[j], (*d)[i] })
	return d
}

func (d *Deck) DrawCard() Card {
	if len(*d) == 0 {
		return Card{Value: -1}
	}
	card := (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	return card
}
