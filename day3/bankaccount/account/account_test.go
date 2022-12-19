package account

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBalance(t *testing.T) {
	acc := Account{balance: 100}

	assert.Equal(t, float64(100), acc.GetBalance())
}

func TestSuccessfulDeposit(t *testing.T) {
	acc := Account{balance: 400}

	(&acc).Deposit(100)

	assert.Equal(t, float64(500), acc.GetBalance())
}

func TestSuccessfulWithdrawal(t *testing.T) {
	acc := Account{balance: 500}

	err := acc.Withdraw(200)

	assert.Nilf(t, err, "Expected withdrawal error to be nil")
	assert.Equal(t, float64(300), acc.GetBalance())
}

func TestWithdrawalError_OnInsufficientBalance(t *testing.T) {
	acc := Account{balance: 150}

	expectedError := &InsufficientBalanceError{balance: 150, amountRequired: 50}

	actualError := acc.Withdraw(200)

	assert.Equal(t, expectedError, actualError)
}
