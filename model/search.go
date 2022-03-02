package model

import (
	"fmt"
	"strings"

	"github.com/kr/pretty"
)

func ListBooks(books []Book) {
	for i := 0; i < len(books); i++ {
		if !books[i].IsDelete {
			//fmt.Println(i, ":", books[i])
			pretty.Println(i, ":", books[i])
		} else {
			fmt.Println("Book is Deleted")
			//fmt.Print("") // You can use just like pass in python
		}
	}
}

func ListBooksWithoutFilter(books []Book) {
	for i := 0; i < len(books); i++ {
		//fmt.Printf("%+v\n", books[i])
		pretty.Println(books[i])
	}
}

// SearchAuthor function search author name in books struct slice
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

// searchBookName function search book name in books struct slice
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

// search function search book sku in books struct slice
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

// SearchAll function is combination of three search functions
func SearchAll(s string, b []Book) ([]Book, []Book, []Book) {
	s = strings.ToLower(s) // TODO: remove after strings.ToLower() in other functions
	authorSlice := SearchAuthor(s, b)
	titleSlice := SearchBookName(s, b)
	skuSlice := SearchSKU(s, b)

	return authorSlice, titleSlice, skuSlice
}

// SearchId function search book id in books struct slice and returns index if not found returns -1
func SearchId(id int, b []Book) int {
	for i := 0; i < len(b); i++ {
		if b[i].ID == id {
			return i
		}
	}
	return -1
}
