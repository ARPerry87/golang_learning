package main

import "fmt"

func main() {
	cards := []string{"Ace of Spades", "Two of Spades", "Three of Spades", newCard()}
	// cards is a slice of strings, which is a list of strings
	cards = append(cards, "Four of Spades") // append is a built-in function that adds an element to the end of a slice
	// this returns a new slice and doesn't alter the original slice
	for i, card := range cards {
		// range is a built-in function that returns the index and value of each element in the slice
		fmt.Println(i, card)
	}
}

func newCard() string {
	// we declare and let the function know that the type that is going to be retunred is a string
	// Whenever this is called we're going to return this value
	return "Five of Diamonds"
}

// var is like in JS and in C, and is used for creating a variable
// string is a type, just like in Ruby and may others, and the type should always be string
// shorthand, := is like declaring var card string = "Ace of Spades"
// with less typing

// Javascript is a dynamically typed language
// Go is a statically typed language, so you have to declare the type of the variable
// in using := you feed it the type you expect, here we fed it a string
// So Go assumes that the card is and should always be a string
// If you try to assign a number to card, it will throw an error

// Only use := inside of a function, outside of a function it needs to be var
// We can also initialize a variable, and then assign a value to it later
// var card string could be called on line 4 and then on line 6 we could do card = "Ace of Spades"

// If you want to reassign a value to a variable, you cannot simply do this:
// card := "Ace of Spades" and then card = "Two of Spades"

// Go has two ways of handling lists

// 1. Arrays, which are fixed lengh list of things
// 2. Slices, which are like arrays but can grow and shrink in size
// Slices are more commonly used in Go, and they are more flexible than arrays
// Both must be declared with a data type and must be of an identical type

// To iterate on every record inside of a slice, we use a for loop
// for i = for the index of the slice, card = for the value of the slice
// so for i, card := range cards means that for the index and card, we want to go through the range of cards
// The range keyword is used to iterate over a slice, array, map, or string
// The range keyword returns two values, the index and the value of the element at that index