package model

import (
	"context"
	"time"
)

type OrderUsecase interface {
	FindAllByCustomerID(ctx context.Context, customerID int64, params OrderParams) ([]Order, int64, error)
	Insert(ctx context.Context, data Order) error
}

type OrderRepository interface {
	FindAllByCustomerID(ctx context.Context, customerID int64, params OrderParams) ([]Order, int64, error)
	Insert(ctx context.Context, data Order) error
}

type Order struct {
	ID              int64     `json:"id"`
	Product         Product   `json:"product"`
	Customer        Customer  `json:"customer"`
	Quantity        int64     `json:"quantity"`
	ShippingAddress string    `json:"shipping_address"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"-"`
}

type OrderParams struct {
	CommonParams
	PaginationParams
}
