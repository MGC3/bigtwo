package game

import (
    "testing"
)

func TestNewPlayedHandSingle(t *testing.T) {
    tenOfDiamonds, _ := NewCard("10", "D")
    singleCardHand := []Card {
        tenOfDiamonds,
    }

    zeroCardHand := []Card {}

    hand, err := NewPlayedHand(singleCardHand)

    if err != nil || hand.Type != single || hand.ToString() != "10D" {
        t.Errorf("Could not create single hand of 10D: %v", hand)
    }

    hand, err = NewPlayedHand(zeroCardHand)
    if err == nil {
        t.Errorf("Created zero card hand: %v", hand)
    }
}

func TestNewPlayedHandPair(t *testing.T) {
    tenOfDiamonds, _ := NewCard("10", "D")
    tenOfHearts, _ := NewCard("10", "H")
    nineOfDiamonds, _ := NewCard("9", "S")

    goodPair := []Card {
        tenOfDiamonds,
        tenOfHearts,
    }

    badPair := []Card {
        tenOfDiamonds,
        nineOfDiamonds,
    }

    hand, err := NewPlayedHand(goodPair)
    if err != nil || hand.Type != pair || hand.ToString() != "10D, 10H" {
        t.Errorf("Could not create good pair: %v, %v", hand, err)
    }

    hand, err = NewPlayedHand(badPair)
    if err == nil {
        t.Errorf("Made bad pair: %v, %v", hand, err)
    }
}

func TestNewPlayedHandTriple(t *testing.T) {
    tenOfDiamonds, _ := NewCard("10", "D")
    tenOfHearts, _ := NewCard("10", "H")
    tenOfSpades, _ := NewCard("10", "S")
    nineOfDiamonds, _ := NewCard("9", "S")

    goodTriple := []Card {
        tenOfHearts,
        tenOfSpades,
        tenOfDiamonds,
    }

    badTriple := []Card {
        tenOfHearts,
        tenOfSpades,
        nineOfDiamonds,
    }

    hand, err := NewPlayedHand(goodTriple)
    if err != nil || hand.Type != triple || hand.ToString() != "10H, 10S, 10D" {
        t.Errorf("Could not create good triple: %v, %v", hand, err)
    }

    hand, err = NewPlayedHand(badTriple)
    if err == nil {
        t.Errorf("Made bad triple: %v, %v", hand, err)
    }
}

func TestNewPlayedHandStraight(t *testing.T) {
    three, _ := NewCard("3", "C")
    four, _ := NewCard("4", "H")
    five, _ := NewCard("5", "C")
    six, _ := NewCard("6", "C")
    seven, _ := NewCard("7", "C")

    goodStraight := []Card {
        seven,
        four,
        five,
        three,
        six,
    }

    hand, err := NewPlayedHand(goodStraight)

    if err != nil || hand.Type != straight || hand.ToString() != "7C, 4H, 5C, 3C, 6C" {
        t.Errorf("Could not create good straight: %v, %v", hand, err)
    }

    eight, _ := NewCard("8", "S")
    badStraight := []Card {
        eight,
        seven,
        six,
        four,
        three,
    }

    hand, err = NewPlayedHand(badStraight)
    if err == nil {
        t.Errorf("Created bad straight: %v", hand)
    }

}
