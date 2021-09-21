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
	a.CustomerId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sql.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
