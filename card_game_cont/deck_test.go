package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	expectedFirstCard := "Ace of Spades"
	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card to be '%s', but got '%s'", expectedFirstCard, d[0])
	}

	expectedLastCard := "Four of Clubs"
	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected last card to be '%s', but got '%s'", expectedLastCard, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Create a new deck and save it to a file
	os.Remove("_decktesting") // Clean up any existing test file

	d := newDeck()
	filename := "_decktesting"
	err := d.saveToFile(filename)
	if err != nil {
		t.Errorf("Error saving to file: %v", err)
	}

	// Read the deck from the file
	loadedDeck, err := newDeckFromFile(filename)
	if err != nil {
		t.Errorf("Error loading from file: %v", err)
	}
	if len(loadedDeck) != 16 {
		t.Errorf("Expected loaded deck length of 16, but got %v", len(loadedDeck))
	}
	expectedFirstCard := "Ace of Spades"
	if loadedDeck[0] != expectedFirstCard {
		t.Errorf("Expected first card to be '%s', but got '%s'", expectedFirstCard, loadedDeck[0])
	}
	expectedLastCard := "Four of Clubs"
	if loadedDeck[len(loadedDeck)-1] != expectedLastCard {
		t.Errorf("Expected last card to be '%s', but got '%s'", expectedLastCard, loadedDeck[len(loadedDeck)-1])
	}
	// Clean up the test file
	err = os.Remove(filename)
	if err != nil {
		t.Errorf("Error removing test file: %v", err)
	}
}
