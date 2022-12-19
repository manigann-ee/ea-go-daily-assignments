package account

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreditCardAccount_GetBalance(t *testing.T) {
	acc := CreditCardAccount{Balance: 100}

	assert.Equal(t, float64(100), acc.GetBalance())
}

func TestCreditCardAccountSuccessful_Deposit(t *testing.T) {
	acc := CreditCardAccount{Balance: 400}

	(&acc).Deposit(100)

	assert.Equal(t, float64(500), acc.GetBalance())
}

func TestCreditCardAccountSuccessful_Withdrawal(t *testing.T) {
	acc := CreditCardAccount{Balance: 500}

	acc.WithDraw(200)

	assert.Equal(t, float64(300), acc.GetBalance())
}
