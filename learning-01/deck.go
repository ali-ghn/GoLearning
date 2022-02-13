package main

import "fmt"

// create a new type of deck which is a slice of string
type deck []string

// a fucntion which prints the deck
func (d deck) printCards() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}
