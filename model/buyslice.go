package model

import "fmt"

// Function takes id , amount and slice of books and do buy book operation
func BuySlice(id, amount int, books []Book) {
	err := books[IDtoIndex(id)].BuyBook(amount)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book Bought")
	}
}
