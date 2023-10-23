package main

import (
	"testing"
)

// TestChangeName tests the changeName function
//
// Tests is based on one case:
//
// 1. The first name is changed
func TestChangeName(t *testing.T) {
	// Declare a variable of type person
	alex := person{firstName: "Ali", lastName: "Ghn"}
	// Declare a variable of type contactInfo and assign it to the field of the person variable
	alex.contactInfo = contactInfo{
		email:   "alighndev@protonmail.com",
		zipCode: 11111111111,
	}
	// change the first name
	alex.changeName("Hey")
	// check if the first name is changed
	if alex.firstName != "Hey" {
		t.Error("The first name has not changed")
	}
}

// TestNotWorkingChangeName tests the notWorkingChangeName function
//
// Tests is based on one case:
//
// 1. The first name is not changed
func TestNotWorkingChangeName(t *testing.T) {
	// Declare a variable of type person
	alex := person{firstName: "Ali", lastName: "Ghn"}
	// Declare a variable of type contactInfo and assign it to the field of the person variable
	alex.contactInfo = contactInfo{
		email:   "alighndev@protonmail.com",
		zipCode: 11111111111,
	}
	// change the first name
	alex.notWorkingChangeName("Hey")
	// check if the first name is changed
	if alex.firstName == "Hey" {
		t.Error("The first name has changed")
	}
}
