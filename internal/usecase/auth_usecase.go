package usecase

import (
	"context"
	"errors"
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

func (u *authUsecase) Login(ctx context.Context, request model.LoginRequest) (token string, err error) {
	logger := log.WithFields(log.Fields{
		"func":    helper.GetFuncName(),
		"request": helper.Dump(request),
	})

	accountData, err := u.accountRepo.FindByUsername(ctx, request.Username)
	if err != nil {
		logger.Errorf("error find account by username, error: %v", err)
		return
	}
	if accountData == nil {
		err = errors.New("incorrect username or password")
		return
	}

	err = helper.ComparePassword(request.Password, accountData.Password)
	if err != nil {
		err = errors.New("incorrect username or password")
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

func (u *authUsecase) Register(ctx context.Context, request model.RegisterRequest) (token string, err error) {
	logger := log.WithFields(log.Fields{
		"func":    helper.GetFuncName(),
		"request": helper.Dump(request),
	})

	customerData := model.Customer{
		Name:  request.Name,
		Email: request.Email,
	}
	newCustomer, err := u.customerRepo.Insert(ctx, customerData)
	if err != nil {
		logger.Errorf("error insert customer, error: %v", err)
		return
	}

	// Hash password
	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		logger.Errorf("error hashing password, error: %v", err)
		return
	}

	accountData := model.Account{
		ID:       newCustomer.ID,
		Username: request.Username,
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
