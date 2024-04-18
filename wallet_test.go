package wallet

import "testing"

func TestWalletDeposit(t *testing.T) {
	t.Parallel()
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	assertBalance(t, wallet, Bitcoin(10))
}

func TestWalletWithdrawWithSufficientFunds(t *testing.T) {
	t.Parallel()
	wallet := Wallet{balance: Bitcoin(20)}
	err := wallet.Withdraw(Bitcoin(10))
	assertNoError(t, err)
	assertBalance(t, wallet, Bitcoin(10))
}

func TestWalletWithdrawWithInsufficientFunds(t *testing.T) {
	t.Parallel()
	wallet := Wallet{balance: Bitcoin(20)}
	err := wallet.Withdraw(Bitcoin(30))
	assertError(t, err, "not enough balance")
	assertBalance(t, wallet, Bitcoin(20))
}

func assertBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != expected {
		t.Errorf("got %f want %f", got, expected)
	}
}

func assertError(t *testing.T, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got.Error() != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
