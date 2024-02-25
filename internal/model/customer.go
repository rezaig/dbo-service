package model

import (
	"context"
	"time"
)

type CustomerUsecase interface {
	FindAll(ctx context.Context) ([]Customer, error)
}

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]Customer, error)
}

type Customer struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
