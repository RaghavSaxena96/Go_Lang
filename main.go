package main

func main() {

	card := newDeck()

	hand, remaining := deal(card, 5)

	hand.print()
	remaining.print()
}
