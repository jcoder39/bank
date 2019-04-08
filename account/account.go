package account

import "github.com/google/uuid"

type Account struct {
	Id      uuid.UUID `json:"id"`
	Balance int       `json:"balance"`
}

func New(balance int) *Account {
	account := new(Account)
	account.Balance = balance
	id, _ := uuid.NewRandom()
	account.Id = id
	return account
}

func (this *Account) Increase(val int) {
	this.Balance += val
}

func (this *Account) Decrease(val int) {
	this.Balance -= val
}
