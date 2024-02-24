package usecase

import "github.com/rezaig/dbo-service/internal/model"

type authUsecase struct {
}

func NewAuthUsecase() model.AuthUsecase {
	return &authUsecase{}
}