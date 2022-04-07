package main

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func(t testing.TB, err error, want error) {
		t.Helper()
		if err == nil {
			t.Fatal("expected an error but got nil")
		}

		if err != want {
			t.Errorf("got %q want %q", err, want)
		}
	}
	assertNoError := func(t testing.TB, got error) {
		t.Helper()

		if got != nil {
			t.Fatal("got error but didn't expect one")
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(200)}
		err := wallet.Withdraw(Bitcoin(100))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(100))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(100)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(500))

		assertBalance(t, wallet, Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
	})
}
