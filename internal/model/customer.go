package model

import (
	"context"
	"time"
)

type CustomerUsecase interface {
	FindAll(ctx context.Context, params CustomerParams) ([]Customer, int64, error)
	FindByID(ctx context.Context, id int64) (*Customer, error)
	Update(ctx context.Context, data Customer, id int64) (*Customer, error)
	Insert(ctx context.Context, data Customer) (*Customer, error)
	Delete(ctx context.Context, id int64) error
}

type CustomerRepository interface {
	FindAll(ctx context.Context, params CustomerParams) ([]Customer, int64, error)
	FindByID(ctx context.Context, id int64) (*Customer, error)
	Update(ctx context.Context, data Customer, id int64) (*Customer, error)
	Insert(ctx context.Context, data Customer) (*Customer, error)
	Delete(ctx context.Context, id int64) error
}

type Customer struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"-"`
}

type CustomerParams struct {
	CommonParams
	PaginationParams
}

// Validate sets default value
func (p *CustomerParams) Validate() error {
	if err := p.PaginationParams.Validate(); err != nil {
		return err
	}
	return nil
}

type CustomerRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
