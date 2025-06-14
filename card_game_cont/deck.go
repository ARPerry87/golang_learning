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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] // we want to return first the slice from the start until the hand size then the handize until the end
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

// To do similar slicing and stepping like in python, you would do something like this:
// fruits= []string{"Apple", "Banana", "Orange", "Pear"}
// fmt.Println(fruits[1:3]) // prints [Banana Orange] - this means it's start inclusive, but up to and not including the end index
// We can also leave out the start or end indexes if we want it to go from the start or until the end
