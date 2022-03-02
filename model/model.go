package model

var BookID int = 0 // Global Book ID stars from zero increment by 1

//Book struct //TODO: For better memory allocation rearrange this struct
type Book struct {
	Author             // Author Struct Embedded
	ID          int    // Id starts from 0 to increment by 1
	BookName    string // Book Name
	ISBN        string // Book ISBN
	PageNumber  int    //
	Price       int    // TODO: change price to float64
	StockAmount int    //
	IsDelete    bool   // Check for book deleted or not

}

// Author struct embeded in Book struct
type Author struct {
	Name       string
	AuthorInfo string
}

// ListBooks function List all books if is not IsDelete set true

// Constructor for Book struct type Book model.Book
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
