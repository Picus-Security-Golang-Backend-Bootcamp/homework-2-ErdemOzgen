package customError

import "errors"

var ErrInSufficentAmount = errors.New("There is no sufficient amount in stocks")
var ErrIDNotInRange = errors.New("Id is not in range")

func InSufficentAmount() error {
	return ErrInSufficentAmount
}

func IDNotInRange() error {
	return ErrIDNotInRange
}
