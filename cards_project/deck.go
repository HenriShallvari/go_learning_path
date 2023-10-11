package main

import "fmt"

//? This file creates a new type of 'deck',
//? defined as a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// using _ tells Go that we don't want to use the index variable in this loop.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// this syntax declares the print() method inside of our file.
// that (d deck) acts as a "receiver", which means that now
// every variable of type "deck" can access this method, which is recalled as such:
// variable.print()

// "d" is the actual variable we are going to use, deck is just our custom type.
// single-letter variable names, although universally hated, are common convention
// in  Go, used to refer to the receiver of a function.
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}
