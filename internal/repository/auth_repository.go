package repository

import (
	"database/sql"
	"github.com/rezaig/dbo-service/internal/model"
)

type authRepository struct {
	dbConn *sql.DB
}

func NewAuthRepository(dbConn *sql.DB) model.AuthRepository {
	return &authRepository{dbConn: dbConn}
}
