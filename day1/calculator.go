package day1

import (
	"errors"
	"math"
)

var errDivisionByZero = errors.New("division by zero")

func add(numberOne, numberTwo int) int {
	return numberOne + numberTwo
}

func subtract(numberOne, numberTwo int) int {
	return numberOne - numberTwo
}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errDivisionByZero
	}
	return x / y, nil
}

func multiply(x, y int) int {
	return x * y
}

func sin(number float64) float64 {
	return math.Sin(number)
}

func cos(number float64) float64 {
	return math.Cos(number)
}

func tan(number float64) float64 {
	return math.Tan(number)
}

func sqrt(number float64) float64 {
	return math.Sqrt(number)
}
