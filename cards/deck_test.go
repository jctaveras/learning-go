package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck of length 52, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected Ace of Clubs as the first card in the deck but got %v", d[0])
	}

	if d[len(d) - 1] != "King of Spades" {
		t.Errorf("Expected King of Spades as the last card in the deck but got %v", d[len(d) - 1])
	}
}

func TestSaveDeckToFileAndGetDeckFromFile(t *testing.T) {
	os.Remove("_testing")

	d := newDeck()

	d.saveToFile("_testing")
	deckFromFile := getDeckFromFile("_testing")

	if len(deckFromFile) != 52 {
		t.Errorf("Expected deck of length 52, but got %v", len(d))
	}

	os.Remove("_testing")
}
