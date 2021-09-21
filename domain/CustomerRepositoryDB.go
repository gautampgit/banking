package domain

import (
	"database/sql"

	"github.com/gautampgit/banking/logger"

	"github.com/gautampgit/banking/errs"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

//
func (db CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var query string
	var rows *sql.Rows
	var err error
	if status == "" {
		query = "select customer_id, name, city,zipcode, date_of_birth, status from customers"
		rows, err = db.client.Query(query)
	} else {
		query = "select customer_id, name, city,zipcode, date_of_birth, status from customers where status = ?"
		rows, err = db.client.Query(query, status)
	}
	customers := make([]Customer, 0)
	if err != nil {
		logger.Error("Error while fetching customers " + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected Database error")
	}
	for rows.Next() {
		var c Customer
		err = rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			logger.Error("Error while fetching customers " + err.Error())
			return nil, errs.NewUnExpectedError("Unexpected error occured")
		}
		customers = append(customers, c)
	}
	//defer db.client.Close()
	return customers, nil
}

//

func (db CustomerRepositoryDB) FindById(id string) (*Customer, *errs.AppError) {
	findIdSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row := db.client.QueryRow(findIdSQL, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while fetching customers " + err.Error())
			return nil, errs.NewNotFoundError("Customer not Found")
		}
		logger.Error("Database Error" + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected Database Error")
	}
	return &c, nil
}

func NewCustomerRepositoryDB(dbClient *sql.DB) CustomerRepositoryDB {

	return CustomerRepositoryDB{dbClient}
}
