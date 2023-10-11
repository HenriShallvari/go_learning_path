package main

import "fmt"

func main() {

	// Slice declaration: notice it wants [] -> datatype of the data it will be storing and
	// the actual data inside of the curly braces
	cards := []string{"Ace Of Diamonds", newCard()}

	// To add an element to the end of a Slice we do
	cards = append(cards, "Six of Spades") // append does not actually modify our slice, instead, it returns a new one.

	//NOTE: every variable in a loop MUST be used, otherwise it generates a syntax error.
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// function type needs to be declared AFTER its name
// except when it is a void function
func newCard() string {
	return "Five of Diamonds"
}
