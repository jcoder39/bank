package account

import "testing"

var account = New(100)

func TestIncrease(t *testing.T) {
	account.Increase(10)
	if account.Balance != 110 {
		t.Error("Expected 110, got ", account.Balance)
	}
}

func TestDecrease(t *testing.T) {
	account.Decrease(10)
	if account.Balance != 100 {
		t.Error("Expected 100, got ", account.Balance)
	}
}