package funding

type Fund struct {
	balance int
}

// A regular function returning a pointer to a fund
func NewFund(initialBalance int) *Fund {
	return &Fund{balance: initialBalance}
}

func (f *Fund) Balance() int {
	return f.balance
}

func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}
