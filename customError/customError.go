package customError

import (
	"errors"
)

var ErrInSufficentAmount = errors.New("there is no sufficient amount in stocks")
var ErrIDNotInRange = errors.New("id is not in range")
var ErrNotInt = errors.New("given value is not int")
var ErrBookIsDeleted = errors.New("book is deleted cant do operation in deleted state")

// Error represents not enough stock amount
func InSufficentAmount() error {
	//fmt.Println(ErrInSufficentAmount)
	return ErrInSufficentAmount
}

// Error represents Id is not in range
func IDNotInRange() error {
	//fmt.Println(ErrIDNotInRange)
	return ErrIDNotInRange
}

// Error represents given value is not int
func NotInt() error {
	//fmt.Println(ErrNotInt)
	return ErrNotInt
}
