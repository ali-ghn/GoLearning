package main

import "fmt"

// contact info for the person
type contactInfo struct {
	email   string
	zipCode int
}

// person struct
type person struct {
	firstName string
	lastName  string
	contactInfo
}

// changeName changes the first name of the person
//
// Using a pointer receiver allows us to modify the original value of the person
//
// Parameters:
// 		newFirstName: the new first name of the person
//
// Returns:
// 		none
//
func (p *person) changeName(newFirstName string) {
	// change the first name of the person
	p.firstName = newFirstName
}

// notWorkingChangeName DOES NOT change the first name of the person as it is not a pointer receiver
//
// Parameters:
// 		newFirstName: the new first name of the person
//
// Returns:
// 		none
//
func (p person) notWorkingChangeName(newFirstName string) {
	// DOES NOT change the first name of the person as it is not a pointer receiver
	p.firstName = newFirstName
}

// print prints the person information details
//
// Returns:
// 		none
func (p person) print() {
	fmt.Println("First name:", p.firstName)
	fmt.Println("Last name:", p.lastName)
	fmt.Println("Email:", p.email)
	fmt.Println("Zip Code:", p.zipCode)

}
