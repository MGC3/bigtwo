package game

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var threeOfClubs, _ = NewCard("3", "C")
var fourOfHearts, _ = NewCard("4", "H")
var fiveOfClubs, _ = NewCard("5", "C")
var sixOfClubs, _ = NewCard("6", "C")
var sevenOfClubs, _ = NewCard("7", "C")
var eightOfSpades, _ = NewCard("8", "S")
var nineOfDiamonds, _ = NewCard("9", "S")
var tenOfHearts, _ = NewCard("10", "H")
var tenOfSpades, _ = NewCard("10", "S")
var tenOfDiamonds, _ = NewCard("10", "D")
var kingOfClubs, _ = NewCard("K", "C")

func assertValidHand(assert *assert.Assertions, hand PlayedHand, err error, expectedType handType, expectedString string) {
    assert.Nil(err)
    assert.Equal(hand.Type, expectedType)
    assert.Equal(hand.ToString(), expectedString)
}

func assertBeats(assert *assert.Assertions, higher PlayedHand, lower PlayedHand) {
    assert.True(higher.Beats(lower))
    assert.False(lower.Beats(higher))
}

func TestNewPlayedHandSingle(t *testing.T) {
    assert := assert.New(t)

    hand, err := NewPlayedHand([]Card{tenOfDiamonds})
    assertValidHand(assert, hand, err, single, "10D")

    lowerHand, err := NewPlayedHand([]Card{fourOfHearts})
    assertValidHand(assert, lowerHand, err, single, "4H")

    assertBeats(assert, hand, lowerHand)

    hand, err = NewPlayedHand([]Card{})
    assert.NotNil(err)
}

func TestNewPlayedHandPair(t *testing.T) {
    assert := assert.New(t)
    tdPair, err := NewPlayedHand([]Card{tenOfDiamonds, tenOfHearts})

    assertValidHand(assert, tdPair, err, pair, "10D, 10H")

    _, err = NewPlayedHand([]Card{nineOfDiamonds, tenOfHearts})
    assert.NotNil(err)
}

func TestNewPlayedHandTriple(t *testing.T) {
    assert := assert.New(t)
    hand, err := NewPlayedHand([]Card {
        tenOfHearts,
        tenOfSpades,
        tenOfDiamonds,
    })

    assertValidHand(assert, hand, err, triple, "10H, 10S, 10D")
    _, err = NewPlayedHand([]Card {
        tenOfHearts,
        tenOfSpades,
        nineOfDiamonds,
    })
    assert.NotNil(err)
}

func TestNewPlayedHandStraight(t *testing.T) {
    assert := assert.New(t)
    hand, err := NewPlayedHand([]Card {
        sevenOfClubs,
        fourOfHearts,
        fiveOfClubs,
        threeOfClubs,
        sixOfClubs,
    })

    assertValidHand(assert, hand, err, straight, "7C, 4H, 5C, 3C, 6C")
    
    _, err = NewPlayedHand([]Card {
        eightOfSpades,
        sevenOfClubs,
        sixOfClubs,
        fourOfHearts,
        threeOfClubs,
    })

    assert.NotNil(err)
    // TODO wrap around straights?
}
