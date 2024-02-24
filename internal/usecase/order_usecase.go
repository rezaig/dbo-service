package usecase

import "github.com/rezaig/dbo-service/internal/model"

type orderUsecase struct {
	orderRepo model.OrderRepository
}

func NewOrderUsecase(orderRepo model.OrderRepository) model.OrderUsecase {
	return &orderUsecase{orderRepo: orderRepo}
}
