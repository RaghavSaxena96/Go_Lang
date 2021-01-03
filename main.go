package main

func main() {

	//card := newDeck()

	//fmt.Println(card.toString())

	//card.saveToFile("my_cards")

	cards := newDeckfromFile("my_cards")

	cards.print()

	//hand, remaining := deal(card, 5)

	//hand.print()
	//remaining.print()

}
