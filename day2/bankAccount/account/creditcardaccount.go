package account

type CreditCardAccount struct {
	balance float64
}

func (cca *CreditCardAccount) GetBalance() float64 {
	return cca.balance
}

func (cca *CreditCardAccount) Deposit(amount float64) {
	cca.balance += amount
}

func (cca *CreditCardAccount) Withdraw(amount float64) {
	cca.balance -= amount
}
