package main

import (
	"fmt"
	"homework-2-ErdemOzgen/banner"
	"homework-2-ErdemOzgen/customError"
	"homework-2-ErdemOzgen/model"
	"homework-2-ErdemOzgen/validator"
	"os"
	"strconv"
	"strings"

	"github.com/kr/pretty"
)

//-----------------GLOBALS----------------
type Book model.Book // type alias

var books []model.Book // books slice for storing initilazed books

//----------------ENDGLOBALS----------------

//-----------------INITIALIZATION----------------
func init() {
	//fmt.Println("init has been started")
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
	//fmt.Println("init has been stopped")
}

//-----------------END INITIALIZATION----------------

//-----------------MAIN----------------
func main() {
	/*
		fmt.Println("main has been started")
		model.ListBooks(books)
		model.BuySlice(1, 200, books)
		books = model.DeleteSlice(1, books)
		model.ListBooks(books)
		e := validator.ValidatorID(2, books)
		fmt.Println(e)
		fmt.Println("main has been stoped")
	*/
	if len(os.Args) == 1 {
		fmt.Println("No args")
		banner.PrintHelp()
		return
	}

	cmds := os.Args[1]

	switch strings.ToLower(cmds) {
	case "buy", "b", "-b", "--buy":
		fmt.Println("Buying...")
		if len(os.Args) <= 2 {
			fmt.Println("Not enough args for buy operation")
			return
		}
		i, _ := strconv.Atoi(os.Args[2])
		err := validator.ValidatorID(i, books)
		if err != nil {
			fmt.Println(err)
			return
		}
		j, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(customError.ErrNotInt)
			return
		}
		model.BuySlice(i, j, books)
		fmt.Println("After Buy operation")
		//model.ListBooks(books) //use for testing list all books
		pretty.Println(books[model.IDtoIndex(i)]) //use for result after buy operation

	case "delete", "d", "--delete", "-d":
		fmt.Println("Deleting...")

		if len(os.Args) <= 2 {
			fmt.Println("Not enough args for delete operation")
			return
		}
		i, _ := strconv.Atoi(os.Args[2])
		err := validator.ValidatorID(i, books)
		if err != nil {
			fmt.Println(err)
			return
		}
		books = model.DeleteSlice(i, books)
		fmt.Println("After Delete operation")
		model.ListBooks(books)

	case "search", "s", "--search", "-s":
		fmt.Println("Searching...")
		if len(os.Args) <= 2 {
			fmt.Println("Not enough args for search operation")
			return
		}
		//i1, i2, i3 := model.SearchAll(os.Args[2], books) // Can be use like this for slice ops
		model.SearchAll(os.Args[2], books)
	case "list", "l", "--list", "-l":
		model.ListBooks(books)

	case "help", "h", "--help", "-h":
		banner.PrintHelp()
	case "version", "v", "--version", "-v":
		banner.PrintBanner()

	default:
		fmt.Println("Unknown command")

	}
}

//-----------------END MAIN----------------n
