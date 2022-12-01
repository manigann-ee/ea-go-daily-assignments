package main

import "errors"

var errDivisionByZero = errors.New("division by zero")

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errDivisionByZero
	}
	return x / y, nil
}