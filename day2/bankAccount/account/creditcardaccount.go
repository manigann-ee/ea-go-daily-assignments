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

func (cca *CreditCardAccount) WithDraw(amount float64) {
	cca.balance -= amount
}

func (cca *CreditCardAccount) CanWithDraw(amount float64) bool {
	return true
}

func (cca *CreditCardAccount) GetIdentifier() string {
	return "CREDIT_CARD_ACCOUNT"
}
