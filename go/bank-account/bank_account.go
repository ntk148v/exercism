package account

import "sync"

// Account reprents a bank account
type Account struct {
	mux     sync.RWMutex
	open    bool
	balance int64
}

// Open creates an Account with a given initial deposit
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{open: true, balance: initialDeposit}
}

// Close closes an Account and returns payout
func (a *Account) Close() (payout int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if !a.open {
		return 0, false
	}
	a.open = false
	if a.balance < 0 {
		return 0, true
	}
	payout = a.balance
	a.balance = 0
	return payout, true
}

// Balance returns an Account's balance
func (a *Account) Balance() (balance int64, ok bool) {
	a.mux.RLock()
	defer a.mux.RUnlock()
	return a.balance, a.open
}

// Deposit add a amount of money to an Account's balance
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if !a.open || (amount < 0 && a.balance+amount < 0) {
		return a.balance, false
	}
	a.balance += amount
	return a.balance, true
}
