/*In Go, methods can only be called on values of a specific type, not on a type itself. 
Therefore, you need to create an instance of the deck type first and 
then call the newDeck() method on that instance.
*/
// main.go
package main

import "fmt"

func main() {
	cards := deck{}          // Create an instance of the deck type
	cards = cards.newDeck()  // Assign the result of newDeck() to cards
	cards.print()            // Print the cards
}

// deck.go
package main

import "fmt"

type deck []string

func (d deck) newDeck() deck {
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d = append(d, value+" of "+suit)
		}
	}

	return d // Return the modified deck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
