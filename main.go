package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	fmt.Printf("\nMy hand is %s and I have %d", hand.toString(), hand.count())
}
