package accounts

import (
	"fmt"
	"learning_go/bank/customers"
)

type SavingsAccount struct {
	Titular                     customers.Titular
	AgencyNumber, NumberMatters int
	balance                     float64
}

func (c *SavingsAccount) Withdraw(value float64) bool {
	canWithdraw := value > 0 && value <= c.balance

	if !canWithdraw {
		fmt.Println("Unable to withdraw")
		return false
	}

	c.balance -= value
	fmt.Println("successful withdrawal")
	return true
}

func (c *SavingsAccount) Deposit(value float64) bool {
	if value <= 0 {
		fmt.Println("Unable to deposit")
		return false
	}

	c.balance += value
	fmt.Println("deposit made successfully")
	return true
}

func (c *SavingsAccount) Transfer(value float64, account *SavingsAccount) bool {
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

func (c *SavingsAccount) getBalance() float64 {
	return c.balance
}
