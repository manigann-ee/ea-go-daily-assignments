package account

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreditCardAccount_GetBalance(t *testing.T) {
	acc := CreditCardAccount{balance: 100}

	assert.Equal(t, float64(100), acc.GetBalance())
}

func TestCreditCardAccountSuccessful_Deposit(t *testing.T) {
	acc := CreditCardAccount{balance: 400}

	(&acc).Deposit(100)

	assert.Equal(t, float64(500), acc.GetBalance())
}

func TestCreditCardAccountSuccessful_Withdrawal(t *testing.T) {
	acc := CreditCardAccount{balance: 500}

	acc.Withdraw(200)

	assert.Equal(t, float64(300), acc.GetBalance())
}
