package wallet

import (
	"errors"
	"sync"
)

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
	mutex   sync.Mutex // Mutex for concurrent access
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if amount > w.balance {
		return errors.New("not enough balance")
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.balance
}
