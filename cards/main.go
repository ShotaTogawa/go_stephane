package main

import (
	"fmt"
)

func main() {

	// cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// fmt.Println(hand, remainingDeck)
	// cards := newDeckFromFile("my_cards")
	// cards.print()
	// cards := newDeck()
	// cards.shuffle()
	// cards.print()

	targetNumbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range targetNumbers {
		if num%2 == 0 {
			fmt.Println(num, " is even")
		} else {
			fmt.Println(num, " v is odd")
		}
	}
}
