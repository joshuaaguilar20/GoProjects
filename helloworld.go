package main 
import "fmt"


func main(){
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string { /* Must be Data Type to Return Function */
	return "Five of Diamonds"
}


