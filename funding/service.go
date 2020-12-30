package funding

import "fmt"

type FundService struct {
	commands chan interface{}
	fund     *Fund
}

func NewFundService(initBalance int) *FundService {
	service := &FundService{
		commands: make(chan interface{}),
		fund:     NewFund(initBalance),
	}
	go service.loop()
	return service
}

func (f *FundService) loop() {
	if f == nil {
		panic("service is nil")
	}
	fmt.Println("Server starts loop...")
	for command := range f.commands {
		switch command.(type) {
		case WithdrawCommand:
			wc := command.(WithdrawCommand)
			f.fund.Withdraw(wc.Amount)
		case BalanceCommand:
			bc := command.(BalanceCommand)
			bc.Response <- f.fund.Balance()
		default:
			panic(fmt.Sprintf("unknown command %v", command))
		}
	}
}

func (f *FundService) Balance() int {
	// fmt.Println("Balance called")
	retChan := make(chan int)
	f.commands <- BalanceCommand{retChan}
	return <-retChan
}

func (f *FundService) Withdraw(amount int) {
	// fmt.Println("Withdraw called")

	f.commands <- WithdrawCommand{amount}
}
