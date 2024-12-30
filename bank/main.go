package main

import (
	"fmt"
	"learning_go/bank/accounts"
	"learning_go/bank/customers"
)

func main() {
	fmt.Println("init")

	luan := customers.Titular{
		Name:       "Luan",
		CPF:        "912.031.233-13",
		Profession: "dev",
	}

	renata := customers.Titular{
		Name:       "Renata",
		CPF:        "401.439.693-52",
		Profession: "dev",
	}

	luanAccount := accounts.CurrentAccount{Titular: luan, AgencyNumber: 123, NumberMatters: 392}
	luanAccount.Deposit(970)

	renataAccount := new(accounts.CurrentAccount)
	renataAccount.Titular = renata
	renataAccount.Deposit(1350)
	renataAccount.AgencyNumber = 321
	renataAccount.NumberMatters = 491

	fmt.Println(luanAccount)
	fmt.Println(renataAccount)

	luanAccount.Withdraw(150)
	renataAccount.Deposit(250)

	fmt.Println(luanAccount)
	fmt.Println(renataAccount)

	luanAccount.Transfer(225, renataAccount)

	fmt.Println(luanAccount)
	fmt.Println(renataAccount)
}
