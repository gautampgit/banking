package domain

import (
	"time"

	"github.com/gautampgit/banking/dto"
)

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionId   string
	AccountId       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}

func NewTransaction(accountId, transactionType string, amount float64) Transaction {
	return Transaction{
		AccountId:       accountId,
		Amount:          amount,
		TransactionType: transactionType,
		TransactionDate: time.Now().Format(TIMEFORMAT),
	}
}
