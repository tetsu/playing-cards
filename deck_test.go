package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 53 {
		t.Errorf("Expected Deck length of 53, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Joker" {
		t.Errorf("Expected last card to be Joker, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)
	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 53 {
		t.Errorf("Expected 53 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove(filename)
}
