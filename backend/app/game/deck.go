package game

import (
	"log"
	"math/rand"
)

type Deck struct {
	Cards []Card
	Index int
}

func NewDeck() Deck {
	ret := Deck{
		Cards: []Card{},
		Index: 0,
	}

	for _, rank := range rankStrings {
		for _, suit := range suitStrings {
			card, err := NewCard(rank, suit)
			if err != nil {
				// Idk this shouldn't happen
				log.Printf("NewDeck couldn't make new card %v\n", err)
				return ret
			}
			ret.Cards = append(ret.Cards, card)
		}
	}

	ret.Shuffle()
	return ret
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Reset() {
	d.Index = 0
	d.Shuffle()
}

func (d *Deck) Deal(n int) []Card {
	// truncate if not enough cards
	if n > d.CardsLeft() {
		n = d.CardsLeft()
	}

	ret := []Card{}
	for ; n > 0; n-- {
		ret = append(ret, d.Cards[d.Index])
		d.Index += 1
	}

	return ret
}

func (d *Deck) CardsLeft() int {
	return len(d.Cards) - d.Index
}
