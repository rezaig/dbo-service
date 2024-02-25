package usecase

import (
	"context"

	"github.com/rezaig/dbo-service/internal/helper"
	"github.com/rezaig/dbo-service/internal/model"
	log "github.com/sirupsen/logrus"
)

type orderUsecase struct {
	orderRepo model.OrderRepository
}

func NewOrderUsecase(orderRepo model.OrderRepository) model.OrderUsecase {
	return &orderUsecase{orderRepo: orderRepo}
}

func (u *orderUsecase) FindAllByCustomerID(ctx context.Context, customerID int64, params model.OrderParams) ([]model.Order, int64, error) {
	results, totalItems, err := u.orderRepo.FindAllByCustomerID(ctx, customerID, params)
	if err != nil {
		log.WithFields(log.Fields{
			"func":       helper.GetFuncName(),
			"customerID": customerID,
			"params":     helper.Dump(params),
		}).Errorf("error find all from repo, error: %v", err)
		return nil, 0, err
	}

	return results, totalItems, nil
}

func (u *orderUsecase) Insert(ctx context.Context, data model.Order) error {
	panic("TODO")
}
