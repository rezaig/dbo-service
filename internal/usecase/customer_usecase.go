package usecase

import "github.com/rezaig/dbo-service/internal/model"

type customerUsecase struct {
}

func NewCustomerUsecase() model.CustomerUsecase {
	return &customerUsecase{}
}
