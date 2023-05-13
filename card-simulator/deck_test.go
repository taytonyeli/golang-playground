package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	fn := "_decktesting"
	os.Remove(fn)

	d := newDeck()
	d.saveToFile(fn)

	ld := newDeckFromFile(fn)

	if len(d) != 20 {
		t.Errorf("Expected deck length of 20, but got %v", len(ld))
	}

	os.Remove(fn)
}
