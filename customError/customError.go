package customError

import "errors"

var ErrInSufficentAmount = errors.New("There is no sufficient amount in stocks")

func InSufficentAmount() error {
	return ErrInSufficentAmount
}
