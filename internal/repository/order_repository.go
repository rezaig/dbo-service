package repository

import (
	"database/sql"
	"github.com/rezaig/dbo-service/internal/model"
)

type orderRepository struct {
	dbConn *sql.DB
}

func NewOrderRepository(dbConn *sql.DB) model.OrderRepository {
	return &orderRepository{dbConn: dbConn}
}
