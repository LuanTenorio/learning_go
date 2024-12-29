package accounts

import (
	"fmt"
	"learning_go/bank/customers"
)

type CurrentAccount struct {
	Titular       customers.Titular
	AgencyNumber  int
	NumberMatters int
	balance       float64
}

func (c *CurrentAccount) Withdraw(value float64) bool {
	canWithdraw := value > 0 && value <= c.balance

	if !canWithdraw {
		fmt.Println("Unable to withdraw")
		return false
	}

	c.balance -= value
	fmt.Println("successful withdrawal")
	return true
}

func (c *CurrentAccount) Deposit(value float64) bool {
	if value <= 0 {
		fmt.Println("Unable to deposit")
		return false
	}

	c.balance += value
	fmt.Println("deposit made successfully")
	return true
}

func (c *CurrentAccount) Transfer(value float64, account *CurrentAccount) bool {
	if c.balance <= 0 {
		return false
	}

	isWithdrawed := c.Withdraw(value)

	if isWithdrawed {
		account.Deposit(value)
		return true
	}

	c.Deposit(value)
	return false
}

func (c *CurrentAccount) getBalance() float64 {
	return c.balance
}
