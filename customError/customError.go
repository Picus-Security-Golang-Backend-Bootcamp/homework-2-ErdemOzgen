package customError

import (
	"errors"
	"fmt"
)

var ErrInSufficentAmount = errors.New("there is no sufficient amount in stocks")
var ErrIDNotInRange = errors.New("id is not in range")
var ErrNotInt = errors.New("given value is not int")

func InSufficentAmount() error {
	fmt.Println(ErrInSufficentAmount)
	return ErrInSufficentAmount
}

func IDNotInRange() error {
	fmt.Println(ErrIDNotInRange)
	return ErrIDNotInRange
}

func NotInt() error {
	fmt.Println(ErrNotInt)
	return ErrNotInt
}
