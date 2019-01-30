package main

func main() {
	cards := newDeck()
	cards.saveToFile("cardFile")

	cards.shuffleDeck()
	cards.print()
}
