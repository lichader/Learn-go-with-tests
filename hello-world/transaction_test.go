package main

import "testing"

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{From: "Alice", To: "Bob", Sum: 100},
		{From: "Bob", To: "Alice", Sum: 25},
	}

	AssertEqual(t, BalanceFor(transactions, "Alice"), -75)
	AssertEqual(t, BalanceFor(transactions, "Bob"), 75)
}

func AssertEqual(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
