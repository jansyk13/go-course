package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length of deck 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expect first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	os.Remove("/tmp/test_deck")

	d := newDeck()

	d.saveToFile("/tmp/test_deck")
	loadedDeck := readFromFile("/tmp/test_deck")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in loadeDeck, but got %v", len(loadedDeck))
	}

	os.Remove("/tmp/test_deck")
}
