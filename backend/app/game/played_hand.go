package game

import (
    "fmt"
)

type handType int

const (
    single handType = iota 
    pair
    triple 

    // Five card hands are below
    straight
    flush
    fullHouse
    fourOfAKind
    straightFlush
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
    if len(cards) == 0 || len(cards) > 5 || len(cards) == 4 {
        return PlayedHand{}, fmt.Errorf("Cannot have hand with %d cards", len(cards))
    }

    if len(cards) == 1 {
        return newHandWithType(cards, single), nil
    }

    if len(cards) == 2 {
        if len(rankCounts(cards)) != 1 {
            return PlayedHand{}, fmt.Errorf("Pair must be of the same rank: %v", cards)
        }
        return newHandWithType(cards, pair), nil
    }

    if len(cards) == 3 {
        if len(rankCounts(cards)) != 1 {
            return PlayedHand{}, fmt.Errorf("Triple must be of the same rank: %v", cards)
        }
        return newHandWithType(cards, triple), nil
    }

    // Handle 5 card hand
    rCounts := rankCounts(cards)
    sCounts := suitCounts(cards)
    highCard := highCard(cards)
    lowCard := lowCard(cards)

    // TODO do I need to handle straights that wrap around?
    // e.g., A2345
    if len(rCounts) == 5 && (highCard.rank - lowCard.rank) == 4 {
        if len(sCounts) == 1 {
            return newHandWithType(cards, straightFlush), nil
        }
        return newHandWithType(cards, straight), nil
    }

    if len(sCounts) == 1 {
        return newHandWithType(cards, flush), nil
    }

    if len(rCounts) == 2 {
        if rCounts[cards[0].rank] == 1 || rCounts[cards[0].rank] == 4 {
            return newHandWithType(cards, fourOfAKind), nil
        }
        return newHandWithType(cards, fullHouse), nil
    }

    return PlayedHand{}, fmt.Errorf("Invalid 5 card hand: %v", cards) 
}

func (p PlayedHand) Count() int {
    return len(p.Cards)
}

func (lhs PlayedHand) Beats(rhs PlayedHand) (bool, error) {
    if lhs.Count() != rhs.Count() {
        return false, fmt.Errorf("Can only compare hands with the same length")
    }

    if lhs.Count () != 5 {
        if lhs.Type != rhs.Type {
            return false, fmt.Errorf("Somehow got <5 card hands with different types? %v, %v", lhs, rhs)
        }
        return highCard(lhs.Cards).GreaterThan(highCard(rhs.Cards)), nil
    }

    if lhs.Type != rhs.Type {
        return lhs.Type > rhs.Type, nil
    }

    if lhs.Type == fullHouse || lhs.Type == fourOfAKind {
        return mostCommonRank(lhs.Cards) > mostCommonRank(rhs.Cards), nil
    }

    // Apparently, rank is determined by the face value of the card in order,
    // and suits are only used as a tiebreaker, but I don't think that's how
    // we used to play
    // https://en.wikipedia.org/wiki/Big_two
    if lhs.Type == flush {
        return lhs.Cards[0].suit > rhs.Cards[0].suit
    }

    // straight and straight flush are ranked by highest card
    return highCard(lhs.Cards).GreaterThan(highCard(rhs.Cards)), nil
}

func (p PlayedHand) ToString() string {
    ret := ""
    for _, c := range p.Cards {
        ret += c.ToString() + ", "
    }
    return ret[:len(ret)-2]
}

//
// Private helper functions
//

// Copies cards into new PlayedHand with type t
// assumes all cards are valid
func newHandWithType(cards []Card, t handType) PlayedHand {
    ret := PlayedHand {
        Cards: make([]Card, len(cards)),
        Type: t,
    }

    for i, c := range cards {
        ret.Cards[i].rank = c.rank
        ret.Cards[i].suit = c.suit
    }

    return ret
}

func rankCounts(cards []Card) map[int]int {
    ret := make(map[int]int)
    for _, card := range cards {
        if _, ok := ret[card.rank]; ok {
            ret[card.rank] += 1
        } else {
            ret[card.rank] = 1
        }
    } 
    return ret
}

func suitCounts(cards []Card) map[int]int {
    ret := make(map[int]int)
    for _, card := range cards {
        if _, ok := ret[card.suit]; ok {
            ret[card.suit] += 1
        } else {
            ret[card.suit] = 1
        }
    } 
    return ret
}

func mostCommonRank(cards []Card) int {
    rCounts := rankCounts(cards)
    mostCommonRank := -1
    numMostCommon := 0

    for rank, count := range(rCounts) {
        if count > numMostCommon {
            numMostCommon = count
            mostCommonRank = rank
        }
    }

    return mostCommonRank
}

func highCard(cards []Card) Card {
    highest := cards[0]
    for _, card := range cards[1:] {
        if card.GreaterThan(highest) {
            highest = card
        }
    }
    return Card{rank: highest.rank, suit: highest.suit}
}

func lowCard(cards []Card) Card {
    lowest := cards[0]
    for _, card := range cards[1:] {
        if lowest.GreaterThan(card) {
            lowest = card
        }
    }
    return Card{rank: lowest.rank, suit: lowest.suit}
}