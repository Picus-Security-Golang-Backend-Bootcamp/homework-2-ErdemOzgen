package model

import (
	"fmt"
	"strings"

	"github.com/kr/pretty"
)

// Function for print all books that not been deleted
func ListBooks(books []Book) {
	for i := 0; i < len(books); i++ {
		if !books[i].IsDelete {
			pretty.Println(i, ":", books[i])
		} else {
			fmt.Println("Book is Deleted")

		}
	}
}

// Function for print all books without comparing anything
func ListBooksWithoutFilter(books []Book) {
	for i := 0; i < len(books); i++ {
		pretty.Println(books[i])
	}
}

// SearchAuthor function search author name in books struct slice
func searchAuthor(s string, b []Book) []Book {
	var result []Book

	for i := 0; i < len(b); i++ {
		a := strings.ToLower(b[i].Author.Name)
		as := strings.Split(a, " ")
		for j := 0; j < len(as); j++ {
			if strings.Contains(s, as[j]) {
				result = append(result, b[i])
			}
		}

	}
	return result
}

// searchBookName function search book name in books struct slice
func searchBookName(s string, b []Book) []Book {
	var result []Book

	for i := 0; i < len(b); i++ {
		if strings.Contains(s, strings.ToLower(b[i].BookName)) {
			result = append(result, b[i])
		}
	}
	return result
}

// search function search book sku in books struct slice
func searchISBN(s string, b []Book) []Book {
	var result []Book

	for i := 0; i < len(b); i++ {
		if strings.Contains(s, strings.ToLower(b[i].ISBN)) {
			result = append(result, b[i])
		}
	}
	return result
}

// Function for print all non empty books
func printSearch(s string, b []Book) {
	if len(b) > 0 {
		pretty.Println(s, ":", b)
	}
}

// SearchAll function is combination of three search functions
func SearchAll(s string, b []Book) ([]Book, []Book, []Book) {
	s = strings.ToLower(s)
	authorSlice := searchAuthor(s, b)
	titleSlice := searchBookName(s, b)
	skuSlice := searchISBN(s, b)
	printSearch("Author Match Slice", authorSlice)
	printSearch("Title Match Slice", titleSlice)
	printSearch("SKU Match Slice", skuSlice)

	return authorSlice, titleSlice, skuSlice
}

// SearchId function search book id in books struct slice and returns index
//if not found returns -1 # i := model.SearchId(2, books) then control i
func SearchId(id int, b []Book) int {
	for i := 0; i < len(b); i++ {
		if b[i].ID == id {
			return i
		}
	}
	return -1
}
