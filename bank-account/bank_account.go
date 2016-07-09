// Package account represents a bank account
package account

import "sync"

// Account is a bank account
type Account struct {
	open    bool
	balance int64
	mutex   sync.Mutex
}

// Open creates a pointer to a new account with a given positive initial deposit
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{open: true, balance: initialDeposit, mutex: sync.Mutex{}}
}

// Close closes the account and returns the current balance
func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.open {
		return 0, false
	}
	currentBalance := a.balance
	a.open = false
	a.balance = 0
	return currentBalance, true
}

// Balance returns the current balance
func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.open {
		return 0, false
	}
	return a.balance, true
}

// Deposit deposit or withdraws a given amount from the account (if the resulting balance is not negative)
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.open || a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
