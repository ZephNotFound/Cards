package main

import "fmt"

func main() {
	cards := []string{"Ace of spades", newCard()}
	cards = append(cards, "King of hearts")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Six of diamonds"
}
