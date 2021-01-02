package main

import "fmt"

func main() {

	var cards string = "Ace of Spades"
	fmt.Println(cards)

	card := deck{"Ace", "Books", "Cars", carder()}

	card.print()
}

func carder() string {

	card := "Ace of Spades"

	return card
}
