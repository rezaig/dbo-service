package repository

import "github.com/rezaig/dbo-service/internal/model"

type orderRepository struct {
}

func NewOrderRepository() model.OrderRepository {
	return &orderRepository{}
}
