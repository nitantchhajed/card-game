package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := newCard()

	//After we initialised a variable with := then we don't have to use := again to assign a new value to it
	//example ->
	//cards = "Five of Diamonds"
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
