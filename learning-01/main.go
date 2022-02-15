package main

import "fmt"

func main() {
	// Create a new deck
	d := newDeck()

	// Save the deck to a file
	d.saveToFile("my_cards")

	// Read the deck from a file
	d, err := newDeckFromFile("my_cards")
	if err != nil {
		panic(err)
	}

	// Shuffle the deck
	d.shuffle()

	// Print the deck
	// d.printCards()

	// Deal the cards
	players, err := d.dealCards(4)
	if err != nil {
		panic(err)
	}

	// Show playes' decks
	for i, player := range players {
		fmt.Println("Player ", i+1, ": ")
		player.printCards()
	}
}
