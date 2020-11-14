package game

import "testing"


func TestNewCard(t *testing.T) {
    c, e := NewCard("3", "C")
    if e != nil || c.Rank != 0 || c.Suit != 0 {
        t.Errorf("Failed to make 3C from string")
    }

    c, e = NewCard("2", "D")
    if e != nil || c.Rank != 12 || c.Suit != 3 {
        t.Errorf("Failed to make 2D from string")
    }

    c, e = NewCard("10", "H")
    if e != nil || c.Rank != 7 || c.Suit != 2 {
        t.Errorf("Failed to make 10H from string")
    }

    c, e = NewCard("H", "D")
    if e == nil || c.Rank != -1 || c.Suit != -1 {
        t.Errorf("Didn't error out on invalid Rank")
    }

    c, e = NewCard("J", "B")
    if e == nil || c.Rank != -1 || c.Suit != -1 {
        t.Errorf("Didn't error out on invalid Suit")
    }
}

func TestCardGreaterThan(t *testing.T) {
    tenOfHearts, _ := NewCard("10", "H")
    tenOfSpades, _ := NewCard("10", "S")
    twoOfDiamonds, _ := NewCard("2", "D")
    
    if !tenOfHearts.GreaterThan(tenOfSpades) || tenOfSpades.GreaterThan(tenOfHearts) {
        t.Errorf("10H !> 10S")
    }

    if !twoOfDiamonds.GreaterThan(tenOfHearts) || tenOfHearts.GreaterThan(twoOfDiamonds) {
        t.Errorf("2D !> 10H")
    }

    if !twoOfDiamonds.GreaterThan(tenOfSpades) || tenOfSpades.GreaterThan(twoOfDiamonds) {
        t.Errorf("2D !> 10S")
    } 
}

func TestCardToString(t *testing.T) {
    tenOfHearts, _ := NewCard("10", "H")
    threeOfClubs, _ := NewCard("3", "C")
    twoOfDiamonds, _ := NewCard("2", "D")
    
    if tenOfHearts.ToString() != "10H" {
        t.Errorf("Bad tostring 10H")
    }

    if threeOfClubs.ToString()  != "3C" {
        t.Errorf("Bad to string 3C")
    }

    if twoOfDiamonds.ToString() != "2D" {
        t.Errorf("Bad to string 2D")
    }
}