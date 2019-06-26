package main 



func main(){
	/* Creating a Slice ( array that grows and shrinks) */
	cards := deck {"Ace of Diamonds", newCard()}
	/* Append Creates new Arrary and Appends Item */
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string { /* Must be Data Type to Return Function */
	return "Five of Diamonds"
}