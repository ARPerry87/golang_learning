package main

func main() {
	cards, err := newDeckFromFile("my_cards")
	if err != nil {
		panic(err)
	}
	cards.print()
}
