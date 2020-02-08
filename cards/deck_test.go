package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedLength := 52
	deckLen := len(d)

	if deckLen != expectedLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedLength, deckLen)
	}
}
