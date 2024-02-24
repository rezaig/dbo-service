package usecase

import (
	"context"

	"github.com/rezaig/dbo-service/internal/model"
)

type customerUsecase struct {
	customerRepo model.CustomerRepository
}

func NewCustomerUsecase(customerRepo model.CustomerRepository) model.CustomerUsecase {
	return &customerUsecase{customerRepo: customerRepo}
}

func (u *customerUsecase) FindAll(ctx context.Context) ([]model.Customer, error) {
	results, err := u.customerRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return results, err
}
