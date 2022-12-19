package quickcash

type QuickCash struct {
	SavingsAccount    Withdrawable
	CreditCardAccount Withdrawable
}

func (qc *QuickCash) getCash(amount float64) (float64, string) {
	if qc.SavingsAccount.CanWithDraw(amount) {
		qc.SavingsAccount.WithDraw(amount)
		return amount, qc.SavingsAccount.GetIdentifier()
	}
	qc.CreditCardAccount.WithDraw(amount)
	return amount, qc.CreditCardAccount.GetIdentifier()
}
