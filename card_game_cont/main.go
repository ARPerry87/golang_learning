package main

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5) // deal the first 5 cards from the deck

	hand.print() // print the dealt hand
	remainingCards.print() // print the remaining cards in the deck
}
