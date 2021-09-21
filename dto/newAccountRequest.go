package dto

import "github.com/gautampgit/banking/errs"

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open account deposit more than 5000")
	}
	if r.AccountType != "saving" || r.AccountType != "checking" {
		return errs.NewValidationError("Account type not savings")
	}
	return nil
}
