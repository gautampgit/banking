package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Gautam", City: "Paradeep", ZipCode: "767687", DateOfBirth: "03/08/95", Status: "1"},
		{"1002", "Sandip", "Hyderabad", "908887", "09/07/96", "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
