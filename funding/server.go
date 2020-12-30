package funding

import "fmt"

type FundServer struct {
	Commands chan interface{}
	fund     *Fund
}

type WithdrawCommand struct {
	Amount int
}
type BalanceCommand struct {
	Response chan int
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		Commands: make(chan interface{}),
		fund:     NewFund(initialBalance),
	}
	go server.loop()
	return server
}

func (f *FundServer) loop() {
	if f == nil {
		panic("service is nil")
	}
	fmt.Println("Server starts loop...")
	for command := range f.Commands {
		switch command.(type) {
		case WithdrawCommand:
			wd := command.(WithdrawCommand)
			f.fund.Withdraw(wd.Amount)

		case BalanceCommand:
			bc := command.(BalanceCommand)
			bc.Response <- f.fund.Balance()
		default:
			panic(fmt.Sprintf("Bad command %v", command))
		}
	}
}
