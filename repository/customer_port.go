package repository

type Customer struct {
	CustomerID  int64  `db:"customer_id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	City        string `db:"city"`
	ZipCode     string `db:"zipcode"`
	Status      int    `db:"status"`
}

type CustomRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
}
