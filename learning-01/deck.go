package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
)

// create a new type of deck which is a slice of string
type deck map[int64]string

// a fucntion which prints the deck
func (d deck) printCards() {
	for i := 0; i < len(d); i++ {
		fmt.Println(i+1, d[int64(i)])
	}
}

// newDeck function which generates a complete set of 52 cards in a deck
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
	// TODO: Create a struct for cards which holds the index of the card and the value of the card
	cardIndex := 0
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards[int64(cardIndex)] = value + " of " + suit
			cardIndex++
		}
	}
	return cards
}

// deal card from the deck with handsize
// func (d deck) dealWithHandSize(handsize int) ([]deck, error) {
// 	// calculate the number of cards to be dealt
// 	if len(d)%handsize != 0 {
// 		return nil, fmt.Errorf("Cannot deal %d cards to %d players", len(d), handsize)
// 	}
// 	start := 0
// 	skip := len(d) / handsize
// 	dealCards := []deck
// 	for i, card := range d {

// 	}
// 	//
// 	return dealCards, nil
// }

// generate a string using the deck
// DEPRECATED
func (d deck) toStringDepricated() string {
	stringDeck := ""
	for i, card := range d {
		if i == int64(len(d)-1) {
			stringDeck += card
		} else {
			stringDeck += card + ","
		}
	}
	return stringDeck
}

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

// convert to a slice of byte
func (d deck) toBytes() []byte {
	return []byte(d.toString())
}

// convert slice to deck
func toDeck(sliceDeck []string) deck {
	mapDeck := deck{}
	for i, card := range sliceDeck {
		mapDeck[int64(i)] = card
	}
	return mapDeck
}

// save the deck to a file
// gets an input for filename of the file
// outputs an error if occures
func (d deck) saveToFile(filename string) error {
	err := ioutil.WriteFile(filename, d.toBytes(), 0666)
	return err
}

// generates a new deck from a file if exists
// gets an input for filename of the file
// ouputs an error if file doesn't exist or an error occures
func newDeckFromFile(filename string) (deck, error) {
	deckInbytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	loadedDeck := strings.Split(string(deckInbytes), ",")
	return deck(toDeck(loadedDeck)), nil
}

// reverses the deck
func (d deck) reverseDeck() {
	for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
		d[int64(i)], d[int64(j)] = d[int64(j)], d[int64(i)]
	}
}

// shuffles the deck
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
