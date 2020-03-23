package main

import "fmt"

func main() {
	fmt.Println("Hi there!")
}

func newCard() string {
	return "Five of diamonds"
}

func variable() {
	// var card string = "Ace of Spades"
	card := "Ace of space"
	card = "Five of Diamonds"
	nCard := newCard()
	fmt.Println(card)
	fmt.Println(nCard)
}

func sliceIteration() {
	// sliceの初期化
	cards := []string{newCard(), newCard()}
	// append
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}
