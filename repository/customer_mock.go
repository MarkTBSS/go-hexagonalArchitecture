package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{
			CustomerID:  1,
			Name:        "John",
			DateOfBirth: "1990-01-01",
			City:        "New York",
			ZipCode:     "10001",
			Status:      1,
		},
		{
			CustomerID:  2,
			Name:        "Jane",
			DateOfBirth: "1990-01-01",
			City:        "New York",
			ZipCode:     "10001",
			Status:      1,
		},
	}
	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == int64(id) {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
