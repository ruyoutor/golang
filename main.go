package main

import "fmt"

func main() {

	cards := newDeck()
	fmt.Println(cards.toString())

	cards.saveToFile("card_file")

	fmt.Println("=== DECK SAVED ===")
	fmt.Println("")

	cards = newDeckFromFile("card_file")
	cards.print()
	fmt.Println("=== DECK OPENEND ===")
}
