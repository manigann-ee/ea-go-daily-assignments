package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdditionOfTwoPositiveNumbers(t *testing.T) {
	result := add(1, 2)
	assert.Equal(t, 3, result, "Expected result to be 3")
}

func TestAdditionOfTwoNegativeNumbers(t *testing.T) {
	result := add(-1, -2)
	assert.Equal(t, -3, result, "Expected result to be -3")
}

func TestSubtractionOfTwoPositiveNumbers(t *testing.T) {
	result := subtract(2, 1)
	assert.Equal(t, 1, result, "Expected result to be 1")
}

func TestDivisionOfTwoNumbers(t *testing.T) {
	result, _ := divide(10, 2)
	assert.Equal(t, float64(5), result, "Expected result to be 5")
}

func TestDivideReturnsErrorWhenNumberIsDividedByZero(t *testing.T) {
	_, err := divide(10, 0)

	assert.EqualError(t, err, "division by zero", "Expected division by zero error")
}

func TestMultiplicationOfTwoNumbers(t *testing.T) {
	result := multiply(3, 3)
	assert.Equal(t, 9, result, "Expected result to be 9")
}

func TestSinOfGivenNumber(t *testing.T) {
	result := sin(1)

	assert.Equal(t, 0.8414709848078965, result, "Expected result to be 0.8414709848078965")
}

func TestCosOfGivenNumber(t *testing.T) {
	result := cos(1)

	assert.Equal(t, 0.5403023058681398, result, "Expected result to be 0.5403023058681398")
}

func TestTanOfGivenNumber(t *testing.T) {
	result := tan(1)

	assert.Equal(t, 1.557407724654902, result, "Expected result to be 1.557407724654902")
}

func TestSquareRootOfGivenNumber(t *testing.T) {
	result := sqrt(4)

	assert.Equal(t, float64(2), result, "Expected result to be 4")
}
