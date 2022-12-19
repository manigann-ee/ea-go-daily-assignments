package account

type SavingsAccount struct {
	Balance float64
}

func (sa *SavingsAccount) GetBalance() float64 {
	return sa.Balance
}

func (sa *SavingsAccount) Deposit(amount float64) {
	sa.Balance += amount
}

func (sa *SavingsAccount) WithDraw(amount float64) {
	sa.Balance -= amount
}

func (sa *SavingsAccount) CanWithDraw(amount float64) bool {
	if sa.Balance >= amount {
		return true
	} else {
		return false
	}
}

func (sa *SavingsAccount) GetIdentifier() string {
	return "PRIMARY_ACCOUNT"
}
