package model

import "context"

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
