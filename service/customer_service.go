package service

import (
	"database/sql"
	"errors"
	"log"

	"github.com/MarkTBSS/go-hexagonalArchitecture/repository"
)

type customerService struct {
	customerRepository repository.ICustomRepository
}

func NewCustomerService(customerRepository repository.ICustomRepository) *customerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

func (s *customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.customerRepository.GetAll()
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	customerResponses := []CustomerResponse{}
	for _, customer := range customers {
		customerResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		customerResponses = append(customerResponses, customerResponse)
	}
	return customerResponses, nil
}

func (s *customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.customerRepository.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}
		log.Println(err)
		return nil, err
	}
	customerRespons := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &customerRespons, nil
}
