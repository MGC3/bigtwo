package game

import (
    "fmt"
)

// rank and suit string arrays are ordered by ascending value
var rankStrings = [...]string {"3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2"}
var suitStrings = [...]string {"C", "S", "H", "D"}

// Cards use characters to represent 
type Card struct {
    rank int 
    suit int 
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
    if lhs.rank == rhs.rank {
        return lhs.suit > rhs.suit
    }

    return lhs.rank > rhs.rank
}

func (c Card) ToString() string {
    if !c.isValid() {
        return "Invalid card"
    }

    return rankStrings[c.rank] + suitStrings[c.suit]
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
    return c.rank >= 0 && c.rank < len(rankStrings) && c.suit >= 0 && c.suit < len(suitStrings)
}