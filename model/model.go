package model

import (
	"fmt"
	"strings"
)

var BookID int64 = 0 // Global Book ID stars from zero increment by 1
type Book struct {
	Author             // Author Struct Embedded
	ID          int64  // Id starts from 0 to increment by 1
	BookName    string // Book Name
	ISBN        string // Book ISBN
	PageNumber  int64  //
	Price       int64  // TODO: change price to float64
	StockAmount int64  //
	IsDelete    bool   // Check for book deleted or not

}
type Author struct {
	Name       string
	AuthorInfo string
}

func ListBooks(books []Book) {
	for i := 0; i < len(books); i++ {
		fmt.Println(books[i])
	}
}

func SearchAuthor(s string, b []Book) []Book {
	var result []Book
	s = strings.ToLower(s) // TODO: remove after SearchAll function fully tested
	for i := 0; i < len(b); i++ {
		if strings.Contains(s, strings.ToLower(b[i].Author.Name)) {
			result = append(result, b[i])
		}
	}
	return result
}
func SearchBookName(s string, b []Book) []Book {
	var result []Book
	s = strings.ToLower(s) // TODO: remove after SearchAll function fully tested
	for i := 0; i < len(b); i++ {
		if strings.Contains(s, strings.ToLower(b[i].BookName)) {
			result = append(result, b[i])
		}
	}
	return result
}
func SearchSKU(s string, b []Book) []Book {
	var result []Book
	s = strings.ToLower(s) // TODO: remove after SearchAll function fully tested
	for i := 0; i < len(b); i++ {
		if strings.Contains(s, strings.ToLower(b[i].ISBN)) {
			result = append(result, b[i])
		}
	}
	return result
}

func SearchAll(s string, b []Book) ([]Book, []Book, []Book) {
	s = strings.ToLower(s) // TODO: remove after strings.ToLower() in other functions
	authorSlice := SearchAuthor(s, b)
	titleSlice := SearchBookName(s, b)
	skuSlice := SearchSKU(s, b)

	return authorSlice, titleSlice, skuSlice
}

func (b *Book) SetDeleted() {
	fmt.Println("Book Delete function", b.IsDelete)
	b.IsDelete = true
	fmt.Println("Book Deleted", b.IsDelete)
}

func NewBook() *Book {
	b := new(Book)
	b.Author.Name = "Erdem"
	b.Author.AuthorInfo = "good author"
	b.ID = BookID
	BookID++
	b.BookName = "Erdem Book"
	b.ISBN = "Erdem ISBN"
	b.PageNumber = 100
	b.Price = 100
	b.StockAmount = 100
	b.IsDelete = false

	return b
}
