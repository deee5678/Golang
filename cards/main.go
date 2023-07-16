package main

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

/*
func main() {

	cards := newDeckFromFile("my_cards")
	cards.print()
}
*/

/*
func main() {
	cards := newDeck()
	// fmt.Println(cards.toString()) -> this prints the deck to strings
	cards.saveToFile("my_cards") //-> this saves the same printed values to file of type string
}
*/

/*
func main() {
	cards := newDeck()
	//cards.print()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}
*/

/* func main() {

	card := "Ace of spades" // := means declaring the variable. this is similar to => var card string = "Ace of spades"
	fmt.Println(card)

}
*/
