package model

import "fmt"

func BuySlice(id, amount int, books []Book) {
	books[IDtoIndex(id)].BuyBook(amount) // TODO: check error if  error return DONOT print
	fmt.Println("Book Bought")           // TODO	: Change to error
}
