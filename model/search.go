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
func searchAuthor(s string, b []Book) []Book {
	var result []Book
	//s = strings.ToLower(s) // TODO: remove after SearchAll function fully tested
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
	//s = strings.ToLower(s) // TODO: remove after SearchAll function fully tested
	for i := 0; i < len(b); i++ {
		if strings.Contains(s, strings.ToLower(b[i].BookName)) {
			result = append(result, b[i])
		}
	}
	return result
}

// search function search book sku in books struct slice
func searchSKU(s string, b []Book) []Book {
	var result []Book
	//s = strings.ToLower(s) // TODO: remove after SearchAll function fully tested
	for i := 0; i < len(b); i++ {
		if strings.Contains(s, strings.ToLower(b[i].ISBN)) {
			result = append(result, b[i])
		}
	}
	return result
}
func printSearch(s string, b []Book) {
	if len(b) > 0 {
		pretty.Println(s, ":", b)
	}
}

// SearchAll function is combination of three search functions
func SearchAll(s string, b []Book) ([]Book, []Book, []Book) {
	s = strings.ToLower(s) // TODO: remove after strings.ToLower() in other functions
	authorSlice := searchAuthor(s, b)
	titleSlice := searchBookName(s, b)
	skuSlice := searchSKU(s, b)
	printSearch("Author Match Slice", authorSlice)
	printSearch("Title Match Slice", titleSlice)
	printSearch("SKU Match Slice", skuSlice)

	return authorSlice, titleSlice, skuSlice
}

// SearchId function search book id in books struct slice and returns index if not found returns -1
/* EXAMPLE
i := model.SearchId(2, books)
if i != -1 {
	fmt.Println("Book is deleted")
}*/
func SearchId(id int, b []Book) int {
	for i := 0; i < len(b); i++ {
		if b[i].ID == id {
			return i
		}
	}
	return -1 //IDNotInRange Error is not implemented yet
}
