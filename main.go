package main

import "fmt"

const deckFile string = "deck_file.txt"

func main() {
	cards := newDeck()
	cards.shuffle()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	fmt.Printf("\nMy hand is %s and I have %d", hand.toString(), hand.count())

	if err := hand.saveToFile(deckFile); err != nil {
		fmt.Printf("Save to fiel error %s", err.Error())
	}
}
