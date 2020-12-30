package funding

import (
	"sync"
	"testing"
)

const WORKERS = 10

func BenchmarkWithdrawlService(b *testing.B) {
	service := NewFundService(100)
	dollarsPerFounder := 100 / WORKERS
	var wg sync.WaitGroup
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < dollarsPerFounder; j++ {
				service.Withdraw(1)
			}
		}()
	}
	wg.Wait()

	balance := service.Balance()
	if balance != 0 {
		b.Error("Should be zero", balance)
	}
}

func BenchmarkWithdrawlServer(b *testing.B) {
	server := NewFundServer(100)
	dollarsPerFounder := 100 / WORKERS
	var wg sync.WaitGroup
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < dollarsPerFounder; j++ {
				server.Commands <- WithdrawCommand{Amount: 1}
			}
		}()
	}
	wg.Wait()
	balanceResponseCh := make(chan int)
	server.Commands <- BalanceCommand{Response: balanceResponseCh}
	balance := <-balanceResponseCh
	if balance != 0 {
		b.Error("Should be zero", balance)
	}
}

// func BenchmarkWithdrawls(b *testing.B) {
// 	if b.N < WORKERS {
// 		return
// 	}
// 	fund := NewFund(b.N)
// 	dollarsPerFounder := b.N / WORKERS
// 	var wg sync.WaitGroup
// 	for i := 0; i < WORKERS; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for j := 0; j < dollarsPerFounder; j++ {
// 				fund.Withdraw(1)
// 			}
// 		}()
// 	}
// 	wg.Wait()
// 	if fund.Balance() != 0 {
// 		b.Error("Should be zero", fund.Balance())
// 	}
// }

// func BenchmarkFund(b *testing.B) {
// 	fund := NewFund(b.N)
// 	for i := 0; i < b.N; i++ {
// 		fund.Withdraw(1)
// 	}
// 	if fund.Balance() != 0 {
// 		b.Error("Should be zero balance:", fund.Balance())
// 	}
// }
