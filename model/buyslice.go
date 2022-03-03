package model

import "fmt"

func BuySlice(id, amount int, books []Book) {
	err := books[IDtoIndex(id)].BuyBook(amount)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book Bought")
	}
}
