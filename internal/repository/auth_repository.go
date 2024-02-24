package repository

import "github.com/rezaig/dbo-service/internal/model"

type authRepository struct {
}

func NewAuthRepository() model.AuthRepository {
	return &authRepository{}
}
