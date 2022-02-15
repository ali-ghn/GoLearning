package main

import (
	"os"
	"strings"
	"testing"
)

// TestNewDeck tests if the new deck is correctly created
//
// Tests are based on three cases:
//
// 1. The deck length is 52
//
// 2. The first card is Ace of Spades
//
// 3. The last card is King of Clubs
func TestNewDeck(t *testing.T) {
	// Create a new deck
	d := newDeck()

	// Assert that the deck is of length
	// 52
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	// Assert that the first card is Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card of the deck to be \"Ace of Spades\", but got %v", d[0])
	}
	// Assert that the last card is King of Clubs
	if d[int64(len(d)-1)] != "King of Clubs" {
		t.Errorf("Expected the last card of the deck to be \"King of Clubs\", but got %v", d[int64(len(d)-1)])
	}
}

// TestToSlice tests if the deck (which is a map) can be successfully converted to a slice
//
// Tests are based on three cases:
//
// 1. The deck length is 52
//
// 2. The first card is Ace of Spades
//
// 3. The last card is King of Clubs
func TestToSlice(t *testing.T) {
	// Create a new deck
	d := newDeck()

	// Convert the deck to a slice
	s := d.toSlice()

	// Assert that the deck is of length
	// 52
	if len(s) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(s))
	}
	// Assert that the first card is Ace of Spades
	if s[0] != "Ace of Spades" {
		t.Errorf("Expected the first card of the deck to be \"Ace of Spades\", but got %v", s[0])
	}
	// Assert that the last card is King of Clubs
	if s[len(s)-1] != "King of Clubs" {
		t.Errorf("Expected the last card of the deck to be \"King of Clubs\", but got %v", s[len(s)-1])
	}
}

// TestToString tests that the deck can be converted to string successfully
//
// Tests are based on four cases:
//
// 1. The first card is Ace of Spades
//
// 2. The last card is King of Clubs
//
// 3. The deck is separated by commas
//
// 4. There are 51 commas in the string
func TestToString(t *testing.T) {
	// Create a new deck
	d := newDeck()

	// Convert the deck to a string
	s := d.toString()

	// Assert that the first card is Ace of Spades
	if s[0:13] != "Ace of Spades" {
		t.Errorf("Expected the first card of the deck to be \"Ace of Spades\", but got %v", s[0])
	}
	// Assert that the last card is King of Clubs
	if s[len(s)-13:] != "King of Clubs" {
		t.Errorf("Expected the last card of the deck to be \"King of Clubs\", but got %v", s[len(s)-13:len(s)])
	}
	// Assert that there are 51 commas in the string
	if strings.Count(s, ",") != 51 {
		t.Errorf("Expected 51 commas in the string, but got %v", strings.Count(s, ","))
	}
}

// TestSaveToFileandNewDeckFromFile tests that the deck can be saved and loaded succesfully
//
// The test has to cleanup the file when starting if it exists and after the test is done
//
// Test filename is _decktest.txt
//
// Tests are based onm five cases:
//
// 1. The deck can be saved to a file
//
// 2. The deck can be loaded from a file
//
// 3. The deck length is 52
//
// 4. The first card is Ace of Spades
//
// 5. The last card is King of Clubs
func TestSaveToFileandNewDeckFromFile(t *testing.T) {
	// Cleanup the file if it exists in directory "cards" with the name _decktest.txt
	filename := "_decktest.txt"
	filename = "cards/" + filename
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}
	filename = "_decktest.txt"
	// Create a new deck
	d := newDeck()

	// Save the deck to a file
	err := d.saveToFile(filename)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Load the deck from a file
	d2, err := newDeckFromFile(filename)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Assert that the deck is of length
	// 52
	if len(d2) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d2))
	}
	// Assert that the first card is Ace of Spades
	if d2[0] != "Ace of Spades" {
		t.Errorf("Expected the first card of the deck to be \"Ace of Spades\", but got %v", d2[0])
	}
	// Assert that the last card is King of Clubs
	if d2[int64(len(d2)-1)] != "King of Clubs" {
		t.Errorf("Expected the last card of the deck to be \"King of Clubs\", but got %v", d2[int64(len(d2)-1)])
	}

	// Cleanup the file when the test is done
	os.Remove("cards/" + filename)
}

// TestDealCards tests that the deck can deal cards correctly
//
// Tests are based on three cases:
//
// 1. If playerQuantity = 4 then each player should have 13 cards
//
// 2. If playerQuantity = 3 then the function should return an error
//
// 3. The players' deck must at least contain one card from the deck
func TestDealCards(t *testing.T) {
	// Create a new deck
	d := newDeck()

	// Deal cards
	players, err := d.dealCards(4)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Assert that each player has 13 cards
	for _, player := range players {
		if len(player) != 13 {
			t.Errorf("Expected each player to have 13 cards, but got %v", len(player))
		}
	}

	// Deal cards
	_, err = d.dealCards(3)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
	// Assert that the players' deck must at least contain one card from the deck
	for _, playerDeck := range players {
		// loop through each card in the player's deck
		for _, card := range playerDeck {
			// Assert that the card is in the deck
			if !d.contains(card) {
				t.Errorf("Expected the player's deck to contain the card %v, but got false", card)
			}
		}
	}
}
