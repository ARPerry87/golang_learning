package main

import (
	"fmt"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// save the deck to the hard drive
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) (deck, error) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err) // if there is an error, we don't want the byteslice and just give us the error
		os.Exit(1)                  // we want to exit the program with a non-zero exit code
	}

	s := strings.Split(string(bs), ",") // we want to convert the byte slice to a string, then split it by the comma
	return deck(s), nil // we want to return a deck type, which is a slice of strings, and a nil error
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

// whenever we start using things that are going to interact with the computers systems or operations we want to use GoLang's built-in packages
// Which we will do for saveToFile and loadFromFile
// so we'd want to go to the GoLang documentation and find the package we want to use
// in this case we'd want to use func WriteFile(filename string, data []byte, perm os.FileMode) error
// and func ReadFile(filename string) ([]byte, error)
// []byte is a slice of bytes, which is what we want to use to write to a file
// byte slices are used often in GoLang to represent data that is not a string
// A byte slice is a slice of bytes, which is to say that if we have a string "Hi there!"
// it would be represented as a slice of bytes like this: []byte{72, 105, 32, 116, 104, 101, 114, 101, 33}
// each letter corresponds to a number in the ASCII table
// so the first letter 'H' is 72, 'i' is 105, and so on

// we're going to use type conversion to convert the deck type to a byte slice
// []byte("Hi there!") is a byte slice that represents the string "Hi there!"

// We pass the error object because we want to know if the error has a value set or if its nil
// if there's a nil we can ignore that object, if it's not we want to know what the error is for newDeckFromFile
