package main

func main() {
	// There are two groups of types in Go:
	//
	// 1. Value types
	// 2. Refrence types
	//
	// Value types are:
	//
	// 1. Numbers
	// 2. Strings
	// 3. Booleans
	// 4. Arrays
	// 5. Structs
	//
	// Refrence types are:
	//
	// 1. Functions
	// 2. Pointers
	// 3. Interfaces
	// 4. Maps
	// 5. Channels
	//
	// Value types are copied when passed to a function or assigned to a variable
	//
	// Refrence types are passed by reference

	// Declare a variable of type person
	ali := person{firstName: "Ali", lastName: "Ghn"}
	// Declare a variable of type contactInfo and assign it to the field of the person variable
	ali.contactInfo = contactInfo{
		email:   "alighndev@protonmail.com",
		zipCode: 11111111111,
	}
	// change the first name
	ali.changeName("Hey")
	ali.print()
}
