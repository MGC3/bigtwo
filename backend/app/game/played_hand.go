package game

import (
    "fmt"
)

/*

TODO

The only valid 5 card hands are straight and above.
How should I handle the different playing modes (singles, pairs, 3 of a kind, hands)?
Currently how it's written is more for a normal poker game, but that's not right
*/

typedef HandType int

const (
    Single HandType = iota 
    Pair
    Triple 

    // Five card hands are below
    Straight
    Flush
    FullHouse
    FourOfAKind
    StraightFlush
    RoyalFlush
)

type PlayedHand struct {
    Cards []Card
    Type handType
}

//
// Public/exported functions
//

// Constructor for a PlayedHand
// Figures out what kind of hand the list of cards is and returns that hand
// If valid, it copies the list of cards into the newly constructed PlayedHand 
// Returns an error if the cards do not create a valid hand
func NewPlayedHand(cards []Card) (PlayedHand, error) {
    if len(cards) == 1 {
        return newHandWithType(cards, highCard), nil 
    }

    if len(cards) == 2 && len(rankCounts(cards)) == 1 {
        return newHandWithType(cards, onePair), nil
    }

    if len(cards) == 3 && len(rankCounts(cards)) == 1 {
        return newHandWithType(cards, threeOfAKind), nil
    }

    if len(cards) == 5 {
        if len(suitCounts(cards)) == 1 {
            return newHandWithType(cards, flush), nil
        }

        rCounts := rankCounts(cards)

        if len(rCounts) == 2 {
            // full house or four of a kind
        }

        if len(rCounts) == 3 {
            return newHandWithType(cards, twoPair), nil
        }

        if len(rCounts) == 4 {
            return newHandWithType(cards, onePair)
        }

    }

    return fmt.Errorf("Invalid hand: %v", cards)
}

func (p *PlayedHand) Count() int {
    return len(p.Cards)
}

func (lhs *PlayedHand) GreaterThan(rhs *PlayedHand) (bool, error) {
}

//
// Private helper functions
//

// Copies cards into new PlayedHand with type t
// assumes all cards are valid
func newHandWithType(cards []Card, t handType) PlayedHand {
    ret := PlayedHand {
        Cards: make([]Card, len(cards)),
        Type: t
    }

    for c := range cards {
        append(ret.Cards, Card{rank: c.rank, suit: c.suit})
    }

    return ret
}

func rankCounts(cards []Card) map[int]int {

}

func suitCounts(cards []Card) map[int]int {

}

func (p *PlayedHand) highCard() *Card {

}