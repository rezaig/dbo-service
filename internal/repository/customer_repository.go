package repository

import (
	"context"
	"database/sql"
	"fmt"

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

func (r *customerRepository) FindAll(ctx context.Context, params model.CustomerParams) ([]model.Customer, int64, error) {
	logger := log.WithFields(log.Fields{
		"func":   helper.GetFuncName(),
		"params": helper.Dump(params),
	})

	selectQ := sq.Select("id", "name", "email", "phone_number", "created_at").
		From("customer")

	if params.Keyword != "" {
		selectQ = selectQ.Where(sq.Like{"name": fmt.Sprintf("%%%s%%", params.Keyword)})
	}

	rows, err := selectQ.
		Limit(uint64(params.PerPage)).
		Offset(uint64(params.GetOffset())).
		RunWith(r.dbConn).
		QueryContext(ctx)
	if err != nil {
		logger.Errorf("error select all, error: %v", err)
		return nil, 0, err
	}

	var results []model.Customer
	for rows.Next() {
		var result model.Customer
		err = rows.Scan(
			&result.ID,
			&result.Name,
			&result.Email,
			&result.PhoneNumber,
			&result.CreatedAt)
		if err != nil {
			logger.Errorf("error scanning, error: %v", err)
			continue
		}
		results = append(results, result)
	}

	selectCountQ := sq.Select("COUNT(id) AS total_items").
		From("customer")
	if params.Keyword != "" {
		selectCountQ = selectCountQ.Where(sq.Like{"name": fmt.Sprintf("%%%s%%", params.Keyword)})
	}
	row := selectCountQ.
		RunWith(r.dbConn).
		QueryRowContext(ctx)
	if err != nil {
		logger.Errorf("error select count, error: %v", err)
		return nil, 0, err
	}
	var totalItems int64
	if err = row.Scan(&totalItems); err != nil {
		logger.Errorf("error scanning, error: %v", err)
		return nil, 0, err
	}

	return results, totalItems, nil
}
