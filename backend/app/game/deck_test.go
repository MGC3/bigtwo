package game

import "testing"


func TestNewCard(t *testing.T) {
    c, e := NewCard("3", "C")
    if e != nil || c.rank != 0 || c.suit != 0 {
        t.Errorf("Failed to make 3C from string")
    }

    c, e = NewCard("2", "D")
    if e != nil || c.rank != 12 || c.suit != 3 {
        t.Errorf("Failed to make 2D from string")
    }

    c, e = NewCard("10", "H")
    if e != nil || c.rank != 7 || c.suit != 2 {
        t.Errorf("Failed to make 10H from string")
    }

    c, e = NewCard("H", "D")
    if e == nil || c.rank != -1 || c.suit != -1 {
        t.Errorf("Didn't error out on invalid rank")
    }

    c, e = NewCard("J", "B")
    if e == nil || c.rank != -1 || c.suit != -1 {
        t.Errorf("Didn't error out on invalid suit")
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