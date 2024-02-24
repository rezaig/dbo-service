package model

import "context"

type CustomerUsecase interface {
	FindAll(ctx context.Context) ([]Customer, error)
}

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]Customer, error)
}

type Customer struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
