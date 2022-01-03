package main

import (
	"fmt"

	"github.com/Pigonhair/go-learning/accounts"
)

func main() {
	account := accounts.NewAccount("kunwoo")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
