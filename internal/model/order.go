package model

import "time"

type OrderUsecase interface {
}

type OrderRepository interface {
}

type Order struct {
	ID              int64     `json:"id"`
	Product         Product   `json:"product"`
	Customer        Customer  `json:"customer"`
	Quantity        int64     `json:"quantity"`
	ShippingAddress string    `json:"shipping_address"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
