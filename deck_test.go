package main

import "testing"

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(cards))
	}
	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card is Ace of Spades, but got %v", cards[0])
	}
	if cards[len(cards)-1] != "King of Hearts" {
		t.Errorf("Expected last card is King of Hearts, but got %v", cards[len(cards)-1])
	}
}
