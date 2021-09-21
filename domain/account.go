package domain

import (
	"github.com/gautampgit/banking/dto"

	"github.com/gautampgit/banking/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Status      string
	Amount      float64
}

func (a Account) ToAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
