package account

import (
	"fmt"
)

type InsufficientBalanceError struct {
	balance        float64
	amountRequired float64
}

func (e *InsufficientBalanceError) Error() string {
	return fmt.Sprintf("Current balance is %d, which is lacking by %d",
		int64(e.balance),
		int64(e.amountRequired))
}
