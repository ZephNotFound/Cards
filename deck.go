package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuite := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValue := []string{"Ace", "two", "three", "four", "five"}

	for _, suit := range cardSuite {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, HandSize int) (deck, deck) {
	return d[:HandSize], d[HandSize:]
}
