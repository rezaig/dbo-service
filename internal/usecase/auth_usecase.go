package usecase

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rezaig/dbo-service/internal/config"
	"github.com/rezaig/dbo-service/internal/helper"
	"github.com/rezaig/dbo-service/internal/model"
	log "github.com/sirupsen/logrus"
)

type authUsecase struct {
	accountRepo  model.AccountRepository
	customerRepo model.CustomerRepository
}

func NewAuthUsecase(accountRepo model.AccountRepository, customerRepo model.CustomerRepository) model.AuthUsecase {
	return &authUsecase{
		accountRepo:  accountRepo,
		customerRepo: customerRepo,
	}
}

func (u *authUsecase) Register(ctx context.Context, req model.RegisterRequest) (token string, err error) {
	logger := log.WithFields(log.Fields{
		"func": helper.GetFuncName(),
		"req":  helper.Dump(req),
	})

	customerData := model.Customer{
		Name:  req.Name,
		Email: req.Email,
	}
	newCustomer, err := u.customerRepo.Insert(ctx, customerData)
	if err != nil {
		logger.Errorf("error insert customer, error: %v", err)
		return
	}

	// Hash password
	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		logger.Errorf("error hashing password, error: %v", err)
		return
	}

	accountData := model.Account{
		ID:       newCustomer.ID,
		Username: req.Username,
		Password: hashedPassword,
	}
	err = u.accountRepo.Insert(ctx, accountData)
	if err != nil {
		logger.Errorf("error insert account, error: %v", err)
		return
	}

	// Generate JWT Token
	timeNowUTC := time.Now().UTC()
	exp := timeNowUTC.Add(config.JWTExp())
	token, err = helper.GenerateJWTToken(jwt.MapClaims{
		"account_id": accountData.ID,
		"exp":        exp.Unix(),
	})
	if err != nil {
		logger.Errorf("error generate token, error: %v", err)
		return
	}

	return
}
