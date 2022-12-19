package account

type PaytmWallet struct {
	Balance float64
}

func (pw *PaytmWallet) GetBalance() float64 {
	return pw.Balance
}

func (pw *PaytmWallet) Deposit(amount float64) {
	pw.Balance += amount
}

func (pw *PaytmWallet) WithDraw(amount float64) {
	pw.Balance -= amount
}

func (pw *PaytmWallet) CanWithDraw(amount float64) bool {
	if pw.Balance >= amount {
		return true
	} else {
		return false
	}
}

func (pw *PaytmWallet) GetIdentifier() string {
	return "PaytmWallet"
}
