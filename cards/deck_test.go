package main

import "testing"

import "os"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedLength := 52
	deckLen := len(d)

	if deckLen != expectedLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedLength, deckLen)
	}

	ex := "Ace of Spades"
	if d[0] != ex {
		t.Errorf("Expected %v but got %v", ex, d[0])
	}

	exEnd := "King of Hearts"
	if d[len(d)-1] != exEnd {
		t.Errorf("Expected %v but got %v", exEnd, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting"
	os.Remove(fileName)

	d := newDeck()
	d.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected a deck of 52 cards but got one of %v", len(loadedDeck))
	}

	os.Remove(fileName)

}
