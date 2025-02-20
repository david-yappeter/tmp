package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (a *BankAccount) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
	fmt.Printf("Deposited: %d | New Balance: %d\n", amount, a.balance)
}

func (a *BankAccount) Withdraw(amount int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	if amount > a.balance {
		fmt.Printf("Withdrawal Failed: %d | Insufficient Funds | Balance: %d\n", amount, a.balance)
		return false
	}

	a.balance -= amount
	fmt.Printf("Withdrew: %d | New Balance: %d\n", amount, a.balance)
	return true
}

func (a *BankAccount) GetBalance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func main() {
	account := &BankAccount{balance: 1000}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(2)

		go func() {
			defer wg.Done()
			account.Deposit(500)
		}()

		go func() {
			defer wg.Done()
			account.Withdraw(300)
		}()
	}

	wg.Wait()
	fmt.Printf("Final Balance: %d\n", account.GetBalance())
}
