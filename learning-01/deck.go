package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
)

// deck is a map of cards
type deck map[int64]string

//printCards prints the cards in the deck
func (d deck) printCards() {
	for i := 0; i < len(d); i++ {
		fmt.Println(i+1, d[int64(i)])
	}
}

// newDeck generates a complete set of 52 cards in a deck
func newDeck() deck {
	// create a new deck of cards
	cards := deck{}
	// create a list of suits
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen ",
		"King",
	}
	cardIndex := 0
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards[int64(cardIndex)] = value + " of " + suit
			cardIndex++
		}
	}
	return cards
}

// dealCards deals the cards beased on the player quantity and the deck size
//
// returns an error in case of five situations:
//
// 1. the deck is empty
//
// 2. the player quantity is invalid
//
// 3. the deck size is invalid
//
// 4. the deck is not divisible by the player quantity
func (d deck) dealCards(playerQuantity int64) ([]deck, error) {
	// check if the deck is empty
	if len(d) == 0 {
		return nil, fmt.Errorf("the deck is empty")
	}
	// check if the player quantity is valid
	if playerQuantity <= 0 {
		return nil, fmt.Errorf("the player quantity is invalid")
	}
	// check if the deck size is valid
	if int64(len(d)) < playerQuantity {
		return nil, fmt.Errorf("the deck size is invalid")
	}
	// check if the deck size is valid
	if int64(len(d))%playerQuantity != 0 {
		return nil, fmt.Errorf("the deck size is invalid")
	}
	// create a new slice of decks
	newDecks := []deck{}
	// shuffle the deck
	newDeck := d
	newDeck.shuffle()
	cardsPerPlayer := int64(len(d)) / playerQuantity
	// loop through the players
	for i := int64(0); i < playerQuantity; i++ {
		// create a new deck
		playerDeck := deck{}
		// loop through the cards per players
		for j := int64(0); j < cardsPerPlayer; j++ {
			// add the card to the player's deck
			playerDeck[j] = newDeck[int64(i)*cardsPerPlayer+j]
		}
		// add the player's deck to the new decks
		newDecks = append(newDecks, playerDeck)
	}

	return newDecks, nil
}

// toSlice converts the deck to a slice of strings
//
// Outputs:
//
// sliceDeck: a slice of strings
func (d deck) toSlice() []string {
	cards := []string{}
	for i := 0; i < len(d); i++ {
		cards = append(cards, d[int64(i)])
	}
	return cards
}

func (d deck) toString() string {
	return strings.Join(d.toSlice(), ",")
}

// toByte converts the deck to a byte array
func (d deck) toBytes() []byte {
	return []byte(d.toString())
}

// toDeck converts a slice of strings to a deck
func toDeck(sliceDeck []string) deck {
	mapDeck := deck{}
	for i, card := range sliceDeck {
		mapDeck[int64(i)] = card
	}
	return mapDeck
}

// saveToFile saves the deck to a file
//
// Inputs:
//
// filename: the name of the file
//
// Outputs:
//
// error: an error if the file couldn't be saved
func (d deck) saveToFile(filename string) error {
	// check if the directory "cards" exists and create it if not
	if _, err := os.Stat("cards"); os.IsNotExist(err) {
		os.Mkdir("cards", 0755)
	}
	// save the deck to a file in "cards" + filename directory
	err := ioutil.WriteFile("cards/"+filename, d.toBytes(), 0644)
	return err
}

// newDeckFromFile generates a new deck from a file if exists
//
// Inputs:
//
// filename: the name of the file
//
// Outputs:
//
// deck: a deck if the file exists
//
// error: an error if the file doesn't exist
func newDeckFromFile(filename string) (deck, error) {
	// check if the directory "cards" exists
	if _, err := os.Stat("cards"); os.IsNotExist(err) {
		return nil, fmt.Errorf("the directory 'cards' doesn't exist")
	}
	// read the file from "cards" + filename directory
	fileBytes, err := ioutil.ReadFile("cards/" + filename)
	if err != nil {
		return nil, err
	}
	loadedDeck := strings.Split(string(fileBytes), ",")
	return deck(toDeck(loadedDeck)), nil
}

// reverseDeck reverses the deck
func (d deck) reverseDeck() {
	for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
		d[int64(i)], d[int64(j)] = d[int64(j)], d[int64(i)]
	}
}

// shuffle shuffles the deck
func (d deck) shuffle() {
	for i := range d {
		newCardPosition, err := rand.Int(rand.Reader, big.NewInt(int64(len(d)-1)))
		if err != nil {
			fmt.Println("CAN NOT SHUFFLE NEW DECK DUE TO LACK OF ENTROPY")
			panic(err)
		}
		d[i], d[int64(newCardPosition.Uint64())] = d[int64(newCardPosition.Uint64())], d[i]
	}
	d.reverseDeck()
	for i := range d {
		newCardPosition, err := rand.Int(rand.Reader, big.NewInt(int64(len(d)-1)))
		if err != nil {
			fmt.Println("CAN NOT SHUFFLE NEW DECK DUE TO LACK OF ENTROPY")
			panic(err)
		}
		d[i], d[int64(newCardPosition.Uint64())] = d[int64(newCardPosition.Uint64())], d[i]
	}
}

// contains checks if a deck contains a card
func (d deck) contains(card string) bool {
	for _, deckCard := range d {
		if deckCard == card {
			return true
		}
	}
	return false
}
