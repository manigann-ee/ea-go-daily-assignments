package quickcash

import (
	"bankAccount/account"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCashFromSavingsAccount(t *testing.T) {
	fpa := &account.SavingsAccount{Balance: 500}
	fsa := &account.CreditCardAccount{}
	accounts := []Withdrawable{
		fpa, fsa,
	}

	fqc := QuickCash{Accounts: accounts}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, fpa.GetIdentifier(), accType)
}

func TestGetCashFromSecondaryAccount(t *testing.T) {

	fpa := &account.SavingsAccount{Balance: 0}
	fsa := &account.CreditCardAccount{Balance: 500}
	accounts := []Withdrawable{
		fpa, fsa,
	}

	fqc := QuickCash{
		Accounts: accounts,
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, fsa.GetIdentifier(), accType)
}

func TestGetCashFromPaytmWallet(t *testing.T) {

	fpa := &account.SavingsAccount{Balance: 0}
	fsa := &account.CreditCardAccount{Balance: 0}
	paytmWallet := &account.PaytmWallet{Balance: 500}
	accounts := []Withdrawable{
		fpa, fsa, paytmWallet,
	}

	fqc := QuickCash{
		Accounts: accounts,
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, paytmWallet.GetIdentifier(), accType)
}
