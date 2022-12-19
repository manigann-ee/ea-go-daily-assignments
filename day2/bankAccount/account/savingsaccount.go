package account

type SavingsAccount struct {
	balance float64
}

func (sa *SavingsAccount) GetBalance() float64 {
	return sa.balance
}

func (sa *SavingsAccount) Deposit(amount float64) {
	sa.balance += amount
}

func (sa *SavingsAccount) WithDraw(amount float64) {
	sa.balance -= amount
}

func (sa *SavingsAccount) CanWithDraw(amount float64) bool {
	return true
}

func (sa *SavingsAccount) GetIdentifier() string {
	return "PRIMARY_ACCOUNT"
}
