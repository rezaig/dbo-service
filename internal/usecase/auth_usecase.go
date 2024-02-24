package usecase

import "github.com/rezaig/dbo-service/internal/model"

type authUsecase struct {
	authRepo model.AuthRepository
}

func NewAuthUsecase(authRepo model.AuthRepository) model.AuthUsecase {
	return &authUsecase{authRepo: authRepo}
}
