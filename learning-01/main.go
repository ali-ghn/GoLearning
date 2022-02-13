package main

func main() {
	// main function which manages cards
	// initial a list called cards which holds a deck of 52 cards
	// cards := newDeck()

	// print cards
	// cards.printCards()

	// deal the deck to 4 players
	// playerQuantity := 4
	// players, err := cards.deal(playerQuantity)
	// if err != nil {
	// 	panic(err)
	// }
	// //  print each player
	// for i, player := range players {
	// 	fmt.Println(i, player)
	// }

	// convert the deck to stirng
	// stringDeck := cards.toString()

	// convert the deck to bytes
	// bytesDeck := cards.toBytes()

	// save the deck to a file
	// cards := newDeck()
	// err := cards.saveToFile("deck01")
	// if err != nil {
	// 	panic(err)
	// }

	// load the deck from a file
	// cards, err := newDeckFromFile("deck01")
	// if err != nil {
	// 	panic(err)
	// }
	// cards.printCards()

	// shuffle the deck
	cards := newDeck()
	cards.shuffle()
	cards.printCards()

}
