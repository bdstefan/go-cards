package main

import "testing"

func TestNewDeck(t * testing.T)  {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	hasAceOfSpades := false
	for _, card := range d {
		if card == "Ace of Spades" {
			hasAceOfSpades = true
			break
		}
	}

	if !hasAceOfSpades {
		t.Error("Expected deck to have Ace of Spades, but it doesn't containt it.")
	}
}