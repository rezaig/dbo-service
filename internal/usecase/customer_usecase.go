package usecase

import (
	"context"
	"errors"

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
		}).Errorf("error find all from repo, error: %v", err)
		return nil, 0, err
	}

	return results, totalItems, nil
}

func (u *customerUsecase) FindByID(ctx context.Context, id int64) (*model.Customer, error) {
	result, err := u.customerRepo.FindByID(ctx, id)
	if err != nil {
		log.WithFields(log.Fields{
			"func": helper.GetFuncName(),
			"id":   id,
		}).Errorf("error find by id from repo, error: %v", err)
		return nil, err
	}

	return result, nil
}

func (u *customerUsecase) Update(ctx context.Context, data model.Customer, id int64) (*model.Customer, error) {
	logger := log.WithFields(log.Fields{
		"func": helper.GetFuncName(),
		"data": helper.Dump(data),
	})

	result, err := u.customerRepo.FindByID(ctx, id)
	if err != nil {
		logger.Errorf("error find by id from repo, error: %v", err)
		return nil, err
	}
	if result == nil {
		return nil, errors.New("data not found")
	}

	updatedData, err := u.customerRepo.Update(ctx, data, id)
	if err != nil {
		logger.Errorf("error update data by id from repo, error: %v", err)
		return nil, err
	}

	return updatedData, nil
}

func (u *customerUsecase) Insert(ctx context.Context, data model.Customer) (*model.Customer, error) {
	insertedData, err := u.customerRepo.Insert(ctx, data)
	if err != nil {
		log.WithFields(log.Fields{
			"func": helper.GetFuncName(),
			"data": helper.Dump(data),
		}).Errorf("error insert data from repo, error: %v", err)
		return nil, err
	}

	return insertedData, nil
}

func (u *customerUsecase) Delete(ctx context.Context, id int64) error {
	panic("TODO")
}
