package main

import (
	"fmt"
	"strings"
)

//------------------STRUCTS------------------
// Arrange field in structs for good memory arrangement
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

//----------------ENDSTRUCTS----------------
//------------------FUNCTIONS AND Methods------------------
// Basic contrructor
func NewErdem() *Book {
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

//------------------ END Contructor------------------
func (b *Book) SetDeleted() {
	fmt.Println("Book Delete function", b.IsDelete)
	b.IsDelete = true
	fmt.Println("Book Deleted", b.IsDelete)
}

func ListBooks(b []Book) {
	for i := 0; i < len(books); i++ {
		fmt.Println(books[i])
	}
}

func searchAuthor(s string, b []Book) []Book {
	var result []Book
	s = strings.ToLower(s) // TODO: remove after SearchAll function fully tested
	for i := 0; i < len(b); i++ {
		if strings.Contains(s, strings.ToLower(b[i].Author.Name)) {
			result = append(result, b[i])
		}
	}
	return result
}
func searchBookName(s string, b []Book) []Book {
	var result []Book
	s = strings.ToLower(s) // TODO: remove after SearchAll function fully tested
	for i := 0; i < len(b); i++ {
		if strings.Contains(s, strings.ToLower(b[i].BookName)) {
			result = append(result, b[i])
		}
	}
	return result
}
func searchSKU(s string, b []Book) []Book {
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
	authorSlice := searchAuthor(s, b)
	titleSlice := searchBookName(s, b)
	skuSlice := searchSKU(s, b)

	return authorSlice, titleSlice, skuSlice
}

//------------------END FUNCTIONS AND Methods------------------

//-----------------GLOBALS----------------
var BookID int64 = 0 // Global variable for book ID
var books []Book

// ----------------END GLOBAL----------------
//-----------------INITIALIZATION----------------
func init() {
	fmt.Println("init")
	e1 := *NewErdem()
	e1.Author.Name = "Erdem1" // testing purpoese
	e1.BookName = "Erdem Book1"
	e1.ISBN = "123456789"
	e2 := *NewErdem()
	e2.Author.Name = "Eren"
	e2.BookName = "Eren at homeland"
	e2.ISBN = "666666"
	e3 := *NewErdem()
	e3.Author.Name = "Elif"
	e3.BookName = "Elif in the disney world"
	e3.ISBN = "888888"
	books = append(books, e1, e2, e3) // add to book slice
	fmt.Println("Books After init", books)
}

//-----------------END INITIALIZATION----------------

//-----------------MAIN----------------

func main() {
	fmt.Println("main")

	//fmt.Println("Books", books)
	fmt.Println("BOOKS")
	for i := 0; i < len(books); i++ {
		fmt.Println("Book", i, ":", books[i].ID)

	}
	ListBooks(books)
	fmt.Println("Search Author")
	e := searchAuthor("Eren", books)
	fmt.Println("e==========>", e)
	f1, f2, f3 := SearchAll("Eren", books)
	fmt.Println("f==========>", f1, f2, f3)

}

//-----------------END MAIN----------------
