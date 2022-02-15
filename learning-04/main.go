package main

import "fmt"

// an interface for bot
type bot interface {
	getGreeting() string
}

// an englishBot struct
type englishBot struct{}

// a spanishBot struct
type spanishBot struct{}

func main() {
	// Interfaces in go are a way of defining a contract
	// between a set of objects and the rest of the program
	// without having to define the behavior of those objects
	// explicitly.
	//
	// For instance in this example we have two bots with different logic
	// in the function getGreeting.
	// Although they both have the same kind of logic in printGreeting.
	//
	// The interface bot defines the contract between the two bots.
	//
	// By implementing the function getGreeting with the same signature (name - parameters - return type)
	// englishBot and spanishBot implement the interface bot.

	// create an englishBot
	eb := englishBot{}
	// create a spanishBot
	sb := spanishBot{}
	// print the english bot greeting
	printGreeting(eb)
	// print the spanish bot greeting
	printGreeting(sb)

}

// getGreeting gets an english bot greeting
//
// Parameters:
// 		none
//
// Returns:
// 		a string containing the english bot greeting
func (englishBot) getGreeting() string {
	// return a string containing the english bot greeting
	return "Hi There!"
}

// getGreeting gets a spanish bot greeting
//
// Parameters:
// 		none
//
// Returns:
// 		a string containing the spanish bot greeting
func (spanishBot) getGreeting() string {
	// return a string containing the spanish bot greeting
	return "Hola!"
}

// printGreeting prints a bot greeting
//
// Parameters:
// 		b: a bot interface
//
// Returns:
// 		none
func printGreeting(b bot) {
	// print the bot greeting
	fmt.Println(b.getGreeting())
}
