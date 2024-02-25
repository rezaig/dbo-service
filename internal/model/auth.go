package model

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	AccountID int64 `json:"account_id"`
	jwt.RegisteredClaims
}

type AuthUsecase interface {
	Login(ctx context.Context, loginRequest LoginRequest) (token string, err error)
	Register(ctx context.Context, registerRequest RegisterRequest) (token string, err error)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
