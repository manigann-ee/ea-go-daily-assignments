package account

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSavingsAccount_GetBalance(t *testing.T) {
	acc := SavingsAccount{balance: 100}

	assert.Equal(t, float64(100), acc.GetBalance())
}

func TestSavingsAccountSuccessful_Deposit(t *testing.T) {
	acc := SavingsAccount{balance: 400}

	(&acc).Deposit(100)

	assert.Equal(t, float64(500), acc.GetBalance())
}

func TestSavingsAccountSuccessful_Withdrawal(t *testing.T) {
	acc := SavingsAccount{balance: 500}

	acc.WithDraw(200)

	assert.Equal(t, float64(300), acc.GetBalance())
}
