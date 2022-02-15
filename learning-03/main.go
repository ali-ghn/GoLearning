package main

import (
	"fmt"
)

func main() {
	// a map of main colours
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	printColours(colours)

	// a map is really similar to a struct bit there are some differences
	//
	// Map
	//
	// 1. All keys must be of the same type
	// 2. All values must be of the same type
	// 3. Keys are indexed and can be iterated over
	// 4. Used to represent a collection of related properties
	// 5. Does not need to know all the keys at compile time
	// 6. Refrence type
	//
	// Struct
	//
	// 1. Values can be of different types
	// 2. Keys do not support indexing
	// 3. You need to know all the field types at compile time
	// 4. Use to represent a "thing" with a set of properties
	// 5. Value type
}

// printColours prints all the colours in the map
//
// Parameters:
// 		colours: the map of colours
//
// Returns:
// 		none
func printColours(colours map[string]string) {
	// loop through the map
	for colour, hex := range colours {
		fmt.Println("Hex code for", colour, "is", hex)
	}
}
