package main

import (
	"os"
	"testing"
)

// go env -w GO111MODULE=auto

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace Spades" {
		t.Errorf("Expected first card of Ace Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King Clubs" {
		t.Errorf("Expected last card of King Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
