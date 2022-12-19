package quickcash

type QuickCash struct {
	Accounts []Withdrawable
}

func (qc *QuickCash) getCash(amount float64) (float64, string) {
	for _, account := range qc.Accounts {
		if account.GetIdentifier() == "PRIMARY_ACCOUNT" && account.CanWithDraw(amount) {
			account.WithDraw(amount)
			return amount, account.GetIdentifier()
		}
		if account.GetIdentifier() == "CREDIT_CARD_ACCOUNT" && account.CanWithDraw(amount) {
			account.WithDraw(amount)
			return amount, account.GetIdentifier()
		}
		if account.GetIdentifier() == "PaytmWallet" && account.CanWithDraw(amount) {
			account.WithDraw(amount)
			return amount, account.GetIdentifier()
		}
	}
	return 0, ""
}
