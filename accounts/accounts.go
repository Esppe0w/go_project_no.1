package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner  string
	blance int
}

// NewAccount
func NewAccount(owner string) *Account {
	account := Account{owner: owner, blance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.blance += amount
}

// Balance of your account
func (a Account) Blance() int {
	return a.blance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {

	if a.blance < amount {
		return errors.New("Can't withdraw no money")
	}
	a.blance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), " 's account.\nHas ", a.Blance())
}
