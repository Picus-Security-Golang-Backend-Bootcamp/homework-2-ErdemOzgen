package validator

import (
	"homework-2-ErdemOzgen/customError"
	"homework-2-ErdemOzgen/model"
	"reflect"
)

func ValidatorInt(a interface{}) error {
	if reflect.ValueOf(a).Kind() == reflect.Int {
		return nil
	} else {
		return customError.NotInt()
	}
}

func ValidatorID(a interface{}, b []model.Book) error {
	if reflect.ValueOf(a).Kind() == reflect.Int {

		e := model.SearchId(a.(int), b)
		if e == -1 {
			return customError.IDNotInRange()
		} else {
			return nil
		}

	} else {
		return customError.NotInt()
	}

}
