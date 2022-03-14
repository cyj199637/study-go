package chapter02

import (
	"fmt"
	"study-go/common"
)

type Account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *Account {
	return &Account{owner: owner, balance: 0}
}

// (a Account) -> copy the receiver
func (a Account) Balance() int {
	return a.balance
}

func (a Account) Owner() string {
	return a.owner
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return common.ErrNoMoney
	}
	a.balance -= amount
	return nil
}

func (a Account) String() string {
	return fmt.Sprint("Account's owner: ", a.owner, "\nCurrent balance: ", a.balance)
}
