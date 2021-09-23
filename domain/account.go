package domain

import (
	"time"

	"github.com/gautampgit/banking/dto"

	"github.com/gautampgit/banking/errs"
)

const TIMEFORMAT = "2006-01-02 15:04:05"

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Status      string
	Amount      float64
}

func (a Account) ToAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{a.AccountId}
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
	FindBy(string) (*Account, *errs.AppError)
}

func NewAccount(customerId, accounttype string, amount float64) Account {
	return Account{
		CustomerId:  customerId,
		OpeningDate: time.Now().Format(TIMEFORMAT),
		AccountType: accounttype,
		Amount:      amount,
		Status:      "1",
	}
}
