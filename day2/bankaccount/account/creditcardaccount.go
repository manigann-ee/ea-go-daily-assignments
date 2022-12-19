package account

type CreditCardAccount struct {
	Balance float64
}

func (cca *CreditCardAccount) GetBalance() float64 {
	return cca.Balance
}

func (cca *CreditCardAccount) Deposit(amount float64) {
	cca.Balance += amount
}

func (cca *CreditCardAccount) WithDraw(amount float64) {
	cca.Balance -= amount
}

func (cca *CreditCardAccount) CanWithDraw(amount float64) bool {
	if cca.Balance >= amount {
		return true
	} else {
		return false
	}
}

func (cca *CreditCardAccount) GetIdentifier() string {
	return "CREDIT_CARD_ACCOUNT"
}
