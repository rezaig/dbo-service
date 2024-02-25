package model

import (
	"time"
)

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"product"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"-"`
}
