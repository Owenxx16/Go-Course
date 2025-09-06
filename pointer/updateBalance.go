package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

// ?
func updateBalance(c *customer, t transaction) error {
	if c.id == t.customerID {
		if c.balance < 0 {
			return errors.New("insufficient funds")
		}
		if t.transactionType != transactionDeposit && t.transactionType != transactionWithdrawal {
			return errors.New("invalid transaction type")
		}
		if t.transactionType == transactionDeposit {
			c.balance += t.amount
		} else if t.transactionType == transactionWithdrawal && c.balance >= t.amount {
			c.balance -= t.amount
		}
	}
	return nil
}
