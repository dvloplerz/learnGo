package repository

import "github.com/jmoiron/sqlx"

type customerRepository struct {
}

func NewCustomerRepository(db *sqlx.DB) customerRepository {
	return customerRepository{}
}
