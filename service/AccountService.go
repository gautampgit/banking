package service

import (
	"github.com/gautampgit/banking/domain"
	"github.com/gautampgit/banking/dto"

	"github.com/gautampgit/banking/errs"
)

const TIMEFORMAT = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)
	newAccount, err := s.repo.Save(a)

	if err != nil {
		return nil, err
	}
	response := newAccount.ToAccountResponseDto()
	return response, nil

}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	//serverside validation to check balance
	if req.IsTransactionWithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient funds")
		}
	}
	t := domain.NewTransaction(req.AccountId, req.TransactionType, req.Amount)
	// t := domain.Transaction{
	// 	AccountId:       req.AccountId,
	// 	Amount:          req.Amount,
	// 	TransactionType: req.TransactionType,
	// 	TransactionDate: time.Now().Format(TIMEFORMAT),
	// }

	transaction, err := s.repo.SaveTransaction(t)
	if err != nil {
		return nil, err
	}
	response := transaction.ToDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
