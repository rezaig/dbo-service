package model

import (
	"context"
	"time"
)

type AccountRepository interface {
	Insert(ctx context.Context, data Account) error
}

type Account struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}
