/*
In Go, you can pass a receiver by reference by using a pointer to the type as the receiver.
This allows you to modify the original value directly within the method.
To do this, you need to use a pointer type for the receiver parameter in the method definition 
and call the method using the address of the variable.
*/
// main.go
package main

import "fmt"

func main() {
	cards := deck{}            // Create an instance of the deck type
	cards.newDeck()            // Call the newDeck() method
	cards.print()              // Print the cards
}

// deck.go
package main

import "fmt"

type deck []string

func (d *deck) newDeck() {
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			*d = append(*d, value+" of "+suit)
		}
	}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}



/*
In this example, the newDeck() method has a receiver parameter of type *deck, which is a pointer to the deck type. 
Within the method, we use *d to dereference the pointer and modify the original value.

When calling the newDeck() method, we pass the address of the cards variable using the & operator: cards.newDeck(). 
This ensures that the method receives a pointer to the original cards variable, allowing it to modify the value directly.

By passing the receiver by reference, any modifications made to *d within the method will affect 
the original value, cards, in the main() function.
*/
