package repository

import (
	"github.com/jmoiron/sqlx"
)

type customRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomRepositoryDB(db *sqlx.DB) *customRepositoryDB {
	return &customRepositoryDB{db}
}
func (r *customRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "SELECT * FROM customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *customRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "SELECT * FROM customers WHERE customer_id = $1"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
