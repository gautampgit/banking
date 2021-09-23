package domain

import (
	"database/sql"
	"strconv"

	"github.com/gautampgit/banking/errs"
	"github.com/gautampgit/banking/logger"
)

type AccountRepositoryDB struct {
	client *sql.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "insert into accounts(customer_id, opening_date, account_type, amount, status) values(?,?,?,?,?)"
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error("Error while adding a new Account " + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected Database Error")
	}
	id, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while retriiving last insert id for  new Account " + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected Database Error")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositoryDB) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	//starting the transaction
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction" + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected database error")
	}

	//inserting bank account transaction
	result, _ := tx.Exec("INSERT INTO transactions (account_id, amount, transaction_type,transaction_date) values(?,?,?,?)", t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	//updating accounts
	if t.IsWithdrawal() {
		_, err = tx.Exec("update accounts set amount = amount - ? where account_id = ?", t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec("update accounts set amount = amount + ? where account_id = ?", t.Amount, t.AccountId)
	}
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction")
		return nil, errs.NewUnExpectedError("Unexpected database error")
	}
	err = tx.Commit()

	if err != nil {
		logger.Error("Error while saving transaction")
		return nil, errs.NewUnExpectedError("Unexpected database error")
	}

	//getting last transactionid
	transactionId, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while retriving transaction id")
		return nil, errs.NewUnExpectedError("Unexpected database error")
	}
	//fetch latest account info
	account, apperr := d.FindBy(t.AccountId)
	if apperr != nil {
		return nil, apperr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDB) FindBy(accountId string) (*Account, *errs.AppError) {
	getAccountQuery := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	row := d.client.QueryRow(getAccountQuery, accountId)
	var account Account
	err := row.Scan(&account.AccountId, &account.CustomerId, &account.OpeningDate, &account.AccountType, &account.Amount)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while fetching customers " + err.Error())
			return nil, errs.NewNotFoundError("Customer not Found")
		}
		logger.Error("Database Error" + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected Database Error")
	}
	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sql.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
