package main

func main() {
	cards := newDeck()
	// var hand deck
	// hand, cards = deal(cards, 4)
	// hand.printDeck()
	cards.shuffle()
	cards.printDeck()
}
