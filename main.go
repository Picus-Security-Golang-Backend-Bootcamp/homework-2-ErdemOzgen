package main

import (
	"fmt"
	"homework-2-ErdemOzgen/model"
)

//-----------------GLOBALS----------------
var BookID int64 = 0 // Global variable for book ID
var books []model.Book

type Book model.Book

//----------------ENDGLOBALS----------------
//------------------Contructor------------------
// Basic Contructor

//------------------ END Contructor------------------

//-----------------INITIALIZATION----------------
func init() {
	fmt.Println("init")
	e1 := *model.NewBook()
	e1.Author.Name = "Erdem Ozgen" // testing purpoese
	e1.BookName = "Erdem Book1"
	e1.ISBN = "1411423410"
	e2 := *model.NewBook()
	e2.Author.Name = "William Shakespeare"
	e2.BookName = "The Tempest"
	e2.ISBN = "1586638491"
	e3 := *model.NewBook()
	e3.Author.Name = "Hall, Franklin"
	e3.BookName = "Glorified Fasting: The Abc of Fasting"
	e3.ISBN = "1684220661"
	books = append(books, model.Book(e1), model.Book(e2), model.Book(e3)) // add to book slice

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
	fmt.Println("Listing All Books")
	model.ListBooks(books)
	fmt.Println("Listing Ends")
	fmt.Println("Search Author")
	e := model.SearchAuthor("Eren", books)
	fmt.Println("e==========>", e)
	f1, f2, f3 := model.SearchAll("Erdem", books)
	fmt.Println("f==========>", f1, f2, f3)

}

//-----------------END MAIN----------------
