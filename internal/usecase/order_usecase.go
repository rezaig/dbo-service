package usecase

import "github.com/rezaig/dbo-service/internal/model"

type orderUsecase struct {
}

func NewOrderUsecase() model.OrderUsecase {
	return &orderUsecase{}
}
