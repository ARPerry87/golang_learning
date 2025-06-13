package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// we used a receiver argument, which is what is before the print()
// by putting what's called a receiver argument in the function
// A receiver argument essentially allows any variable of type 'deck'
// access to the print method
// Sets up methods on variables we create

// The d is the variable that gets passed into the function
// While we use the word deck to define the type for the variable
// don't use this or self in go, always refer to the receiver by what it actually is
// receiver arguments should only be one or two characters that matches the type of the receiver
