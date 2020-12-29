package game

import (
	"fmt"
	"log"
)

// rank and suit string arrays are ordered by ascending value
var rankStrings = [...]string{"3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2"}
var suitStrings = [...]string{"C", "S", "H", "D"}

type Card struct {
	Rank int
	Suit int
}

type JsonCard struct {
	Rank string `json:"rank"`
	Suit string `json:"suit"`
}

//
// Public/exported functions
//

// Constructor from a card from rank/suit as string values
// String values must match those in rankStrings/suitStrings exactly
// Returns an error if rank or suit is invalid
func NewCard(rank string, suit string) (Card, error) {
	rankInt, err := rankIntFromString(rank)
	if err != nil {
		return invalidCard(), err
	}

	suitInt, err := suitIntFromString(suit)
	if err != nil {
		return invalidCard(), err
	}

	return Card{rankInt, suitInt}, nil
}

// Returns true if c > other, false otherwise
// Assumes both c and other are valid cards
func (lhs Card) GreaterThan(rhs Card) bool {
	if lhs.Rank == rhs.Rank {
		return lhs.Suit > rhs.Suit
	}

	return lhs.Rank > rhs.Rank
}

func (c Card) ToString() string {
	if !c.isValid() {
		return "Invalid card"
	}

	return rankStrings[c.Rank] + suitStrings[c.Suit]
}

func (c Card) ToJsonCard() JsonCard {
	return JsonCard{
		Rank: rankStrings[c.Rank],
		Suit: suitStrings[c.Suit],
	}
}

func (c Card) Equals(other Card) bool {
	return c.Rank == other.Rank && c.Suit == other.Suit
}

func (j JsonCard) ToCard() (Card, error) {
	return NewCard(j.Rank, j.Suit)
}

func CardListToJson(cards []Card) []JsonCard {
	ret := []JsonCard{}
	for _, c := range cards {
		ret = append(ret, c.ToJsonCard())
	}
	return ret
}

func CardListFromJson(jsonCards []JsonCard) ([]Card, error) {
	ret := []Card{}
	for _, j := range jsonCards {
		card, err := j.ToCard()
		if err != nil {
			log.Printf("CardListFromJson failed %v\n", err)
			return ret, err
		}
		ret = append(ret, card)
	}
	return ret, nil
}

func CardInList(card Card, list []Card) bool {
	for _, c1 := range list {
		if c1.Equals(card) {
			return true
		}
	}

	return false
}

// Creates a new slice without cards in toRemove, returns an error
// and the original hand if a card in toRemove is not in hand
// TODO do this in place, or not necessary?
func RemoveCardsFromList(hand, toRemove []Card) ([]Card, error) {
	// Validate that all cards in toRemove are in hand
	for _, card := range toRemove {
		if !CardInList(card, hand) {
			return hand, fmt.Errorf("error card %v not in hand %v\n", card, hand)
		}
	}

	// Actually remove cards in second pass
	ret := []Card{}
	for _, card := range hand {
		if !CardInList(card, toRemove) {
			ret = append(ret, card)
		}
	}

	return ret, nil
}

//
// Private Helper functions
//
func rankIntFromString(rank string) (int, error) {
	for index := range rankStrings {
		if rankStrings[index] == rank {
			return index, nil
		}
	}
	return -1, fmt.Errorf("Could not find rank for %s", rank)
}

func suitIntFromString(suit string) (int, error) {
	for index := range suitStrings {
		if suitStrings[index] == suit {
			return index, nil
		}
	}
	return -1, fmt.Errorf("Could not find suit for %s", suit)
}

func invalidCard() Card {
	return Card{-1, -1}
}

func (c Card) isValid() bool {
	return c.Rank >= 0 && c.Rank < len(rankStrings) && c.Suit >= 0 && c.Suit < len(suitStrings)
}
