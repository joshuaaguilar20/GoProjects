package main 
import "fmt"


/* Types Allow Custom Fuctions To work with type like class methods  
	Slice of Strings

*/
type deck [] string 


func(d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	}
}; 


