package main

import "fmt"

func main() {

	// Slice declaration: notice it wants [] -> datatype of the data it will be storing and
	// the actual data inside of the curly braces
	cards := []string{"Ace Of Diamonds", newCard()}

	// To add an element to the end of a Slice we do
	cards = append(cards, "Six of Spades")

	fmt.Println(cards)
}

// function type needs to be declared AFTER its name
// except when it is a void function
func newCard() string {
	return "Five of Diamonds"
}
