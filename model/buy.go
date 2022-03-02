package model

import "fmt"

func Buy(id, amount int, books []Book) {
	books[IDtoIndex(id)].BuyBook(amount)
	fmt.Println("Book Bought")
}
