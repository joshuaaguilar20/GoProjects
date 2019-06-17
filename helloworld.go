package main 
import "fmt"


func main(){
	/* Creating a Slice ( array that grows and shrinks) */
	cards := []string{"Ace of Diamonds", newCard()}
	/* Append Creates new Arrary and Appends Item */
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		/* Loops over Slice to Print each Card */
		fmt.Println(i, card)
	}
}

func newCard() string { /* Must be Data Type to Return Function */
	return "Five of Diamonds"
}


