package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("Withdraw failed, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		// return fmt.Errorf("Insufficient balance")
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w Wallet) Balance() Bitcoin {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	return w.balance
}

type Stringer interface {
	String() string
}
