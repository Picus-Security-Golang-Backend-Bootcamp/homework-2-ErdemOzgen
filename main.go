package main

import (
	"fmt"
	"homework-2-ErdemOzgen/model"
)

//-----------------GLOBALS----------------
type Book model.Book // type alias

var books []model.Book // books slice for storing initilazed books

//----------------ENDGLOBALS----------------

//-----------------INITIALIZATION----------------
func init() {
	fmt.Println("init has been started")
	e1 := *model.NewBook()
	e1.Author.Name = "Erdem" // testing purpoese
	e1.BookName = "Erdem Book1"
	e1.ISBN = "1411423410"
	e1.IsDelete = false
	e2 := *model.NewBook()
	e2.Author.Name = "William Shakespeare"
	e2.BookName = "The Tempest"
	e2.ISBN = "1586638491"
	e2.IsDelete = false
	e3 := *model.NewBook()
	e3.Author.Name = "Hall"
	e3.BookName = "Glorified Fasting: The Abc of Fasting"
	e3.ISBN = "1684220661"
	e3.IsDelete = false
	books = append(books, model.Book(e1), model.Book(e2), model.Book(e3)) // add to book slice
	fmt.Println("init has been stopped")
}

//-----------------END INITIALIZATION----------------

//-----------------MAIN----------------
func main() {
	fmt.Println("main has been started")
	model.ListBooks(books)
	model.Buy(1, 200, books)
	books[0].Delete(1)
	model.ListBooks(books)
	fmt.Println("main has been stoped")

}

//-----------------END MAIN----------------n
