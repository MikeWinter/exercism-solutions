package account

import "sync"

type Money int64

// Opens a new Account. Returns nil if initialDeposit is negative.
func Open(initialDeposit Money) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit}
}

type Account struct {
	balance Money
	closed bool
	mux sync.Mutex
}

// Adds an amount to the Account, returning the new balance. If amount is
// negative, then a withdrawal occurs. If this would result in a negative
// balance or the Account has been closed then ok is false; otherwise true.
func (acc *Account) Deposit(amount Money) (balance Money, ok bool) {
	acc.mux.Lock()
	defer acc.mux.Unlock()

	if acc.closed {
		return 0, false
	}
	if amount < 0 && -amount > acc.balance {
		return acc.balance, false
	}

	acc.balance += amount
	return acc.balance, true
}

// Returns the current balance on an open Account. If the account is closed
// ok is false; otherwise true.
func (acc *Account) Balance() (balance Money, ok bool) {
	acc.mux.Lock()
	defer acc.mux.Unlock()

	if acc.closed {
		return 0, false
	}
	return acc.balance, true
}

// Closes an Account returning the balance in payout. If the account has
// already been closed then ok is false; otherwise true.
func (acc *Account) Close() (payout Money, ok bool) {
	acc.mux.Lock()
	defer acc.mux.Unlock()

	if acc.closed {
		return 0, false
	}

	acc.closed = true
	return acc.balance, true
}
