package main

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.rint()
	// remainingCards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
