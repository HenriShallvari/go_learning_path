package main // ? declares an executable package
import "fmt"

// ! mandatory with package main
func main() {

	// type-inferred variable -> the compiler figures out the type of this variable
	// based on the right side of the expression. In this case, we will have a string
	// because the newCard() function returns a string

	// the long way of declaring a variable would be:
	// var card string = newCard() -> this way we inform the compiler that we want a string
	card := newCard()

	fmt.Println(card)
}

// function type needs to be declared AFTER its name
// except when it is a void function
func newCard() string {
	return "Five of Diamonds"
}
