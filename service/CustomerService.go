package service

import (
	"github.com/gautampgit/banking/dto"

	"github.com/gautampgit/banking/errs"

	"github.com/gautampgit/banking/domain"
)

type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerById(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	switch status {
	case "active":
		status = "1"
	case "inactive":
		status = "0"
	default:
		status = ""
	}
	customers := make([]dto.CustomerResponse, 0)
	c, err := d.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	for _, v := range c {
		customers = append(customers, v.ToDo())
	}
	return customers, nil
}

func (d DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := d.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDo()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
