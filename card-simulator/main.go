package main

func main() {
	cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())

	// cards.saveToFile("sample_cards")

	// cards := newDeckFromFile("sample_cardss")
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	cards.print()

	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Spades"
}
