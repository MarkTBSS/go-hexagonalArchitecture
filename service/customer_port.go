package service

//DTO
type CustomerResponse struct {
	CustomerID int64  `json:"customer_id" xml:"customer_id"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
}

type CustomService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(id int) (*CustomerResponse, error)
}
