package main

import (
	"github.com/tazhibayda/wallet"
)

func main() {
	w := wallet.Wallet{}

	w.Deposit(100.00)
	println(int(w.Balance()))
	err := w.Withdraw(100.00)
	if err != nil {
		println(err.Error())
		return
	}
	println(int(w.Balance()))
	err = w.Withdraw(100.00)
	if err != nil {
		println(err.Error())
		return
	}
}
