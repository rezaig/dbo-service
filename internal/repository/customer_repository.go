package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/rezaig/dbo-service/internal/helper"
	"github.com/rezaig/dbo-service/internal/model"
	log "github.com/sirupsen/logrus"
)

type customerRepository struct {
	dbConn *sql.DB
}

func NewCustomerRepository(dbConn *sql.DB) model.CustomerRepository {
	return &customerRepository{dbConn: dbConn}
}

func (r *customerRepository) FindAll(ctx context.Context) ([]model.Customer, error) {
	logger := log.WithFields(log.Fields{
		"ctx":  helper.Dump(ctx),
		"func": helper.GetFuncName(),
	})

	rows, err := sq.Select("id", "name").
		From("customer").
		RunWith(r.dbConn).
		QueryContext(ctx)
	if err != nil {
		logger.Errorf("error select all, error: %v", err)
		return nil, err
	}

	var results []model.Customer
	for rows.Next() {
		var result model.Customer
		err = rows.Scan(&result.ID, &result.Name)
		if err != nil {
			logger.Errorf("error scanning, error: %v", err)
			continue
		}
		results = append(results, result)
	}

	return results, nil
}
