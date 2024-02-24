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

func (u *customerUsecase) FindAll(ctx context.Context) ([]model.Customer, error) {
	logger := log.WithFields(log.Fields{
		"func": helper.GetFuncName(),
	})

	results, err := u.customerRepo.FindAll(ctx)
	if err != nil {
		logger.Errorf("error find all, error: %v", err)
		return nil, err
	}

	return results, err
}
