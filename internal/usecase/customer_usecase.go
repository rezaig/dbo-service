package usecase

import (
	"context"

	"github.com/rezaig/dbo-service/internal/helper"
	"github.com/rezaig/dbo-service/internal/model"
	log "github.com/sirupsen/logrus"
)

type customerUsecase struct {
	customerRepo model.CustomerRepository
}

func NewCustomerUsecase(customerRepo model.CustomerRepository) model.CustomerUsecase {
	return &customerUsecase{customerRepo: customerRepo}
}

func (u *customerUsecase) FindAll(ctx context.Context, params model.CustomerParams) ([]model.Customer, int64, error) {
	results, totalItems, err := u.customerRepo.FindAll(ctx, params)
	if err != nil {
		log.WithFields(log.Fields{
			"func":   helper.GetFuncName(),
			"params": helper.Dump(params),
		}).Errorf("error find all, error: %v", err)
		return nil, 0, err
	}

	return results, totalItems, nil
}
