package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with less balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		error := wallet.Withdraw(Bitcoin(10))
		if error != nil {
			t.Error("got an error but didn't want one")
		}

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with exact same balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		error := wallet.Withdraw(Bitcoin(20))
		if error != nil {
			t.Error("got an error but didn't want one")
		}
		assertBalance(t, wallet, Bitcoin(1))
	})

	t.Run("withdraw with more balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted an error but got nothing")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
