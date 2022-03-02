package model

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
