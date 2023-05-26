package main

func main() {
	cards := newDeck()
	//cards.print()

	// top, bot := deal(cards, 4)
	// top.print()
	// fmt.Println()
	// bot.print()

	// cards.shuffle()
	// cards.print()

	cards.saveToFile("my_cards")
	loadedCards := readFromFile("my_cards")
	loadedCards.print()
}
