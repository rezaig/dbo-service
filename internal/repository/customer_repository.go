package repository

import "github.com/rezaig/dbo-service/internal/model"

type customerRepository struct {
}

func NewCustomerRepository() model.CustomerRepository {
	return &customerRepository{}
}
