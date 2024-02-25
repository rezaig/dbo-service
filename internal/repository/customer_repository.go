package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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
		logger.Errorf("error run query select all, error: %v", err)
		return nil, 0, err
	}

	var results []model.Customer
	for rows.Next() {
		var result model.Customer
		var phoneNumber sql.NullString
		err = rows.Scan(
			&result.ID,
			&result.Name,
			&result.Email,
			&phoneNumber,
			&result.CreatedAt)
		if err != nil {
			logger.Errorf("error scanning query select all, error: %v", err)
			continue
		}
		result.PhoneNumber = phoneNumber.String
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
		logger.Errorf("error run query select count, error: %v", err)
		return nil, 0, err
	}
	var totalItems int64
	if err = row.Scan(&totalItems); err != nil {
		logger.Errorf("error scanning query select count, error: %v", err)
		return nil, 0, err
	}

	return results, totalItems, nil
}

func (r *customerRepository) FindByID(ctx context.Context, id int64) (*model.Customer, error) {
	row := sq.Select("id", "name", "email", "phone_number", "created_at").
		From("customer").
		Where(sq.Eq{"id": id}).
		RunWith(r.dbConn).QueryRowContext(ctx)

	result := new(model.Customer)
	err := row.Scan(
		&result.ID,
		&result.Name,
		&result.Email,
		&result.PhoneNumber,
		&result.CreatedAt)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.WithFields(log.Fields{
			"func": helper.GetFuncName(),
			"id":   id,
		}).Errorf("error scan query select by id, error: %v", err)
		return nil, err
	}

	return result, nil
}

func (r *customerRepository) Update(ctx context.Context, data model.Customer, id int64) (*model.Customer, error) {
	logger := log.WithFields(log.Fields{
		"func": helper.GetFuncName(),
		"data": helper.Dump(data),
	})

	timeNow := time.Now().UTC()
	_, err := sq.Update("customer").
		Set("name", data.Name).
		Set("email", data.Email).
		Set("phone_number", data.PhoneNumber).
		Set("updated_at", timeNow).
		Where(sq.Eq{"id": id}).
		RunWith(r.dbConn).
		ExecContext(ctx)
	if err != nil {
		logger.Errorf("error exec query update, error: %v", err)
		return nil, err
	}

	data.ID = id
	data.UpdatedAt = timeNow

	return &data, nil
}

func (r *customerRepository) Insert(ctx context.Context, data model.Customer) (*model.Customer, error) {
	logger := log.WithFields(log.Fields{
		"func": helper.GetFuncName(),
		"data": helper.Dump(data),
	})

	var phoneNumber sql.NullString
	if data.PhoneNumber != "" {
		phoneNumber = sql.NullString{String: data.PhoneNumber, Valid: true}
	}

	timeNow := time.Now().UTC()
	_, err := sq.Insert("customer").
		Columns("name", "email", "phone_number", "created_at").
		Values(data.Name, data.Email, phoneNumber, timeNow).
		RunWith(r.dbConn).
		ExecContext(ctx)
	if err != nil {
		logger.Errorf("error exec query insert, error: %v", err)
		return nil, err
	}

	var insertedID int64
	err = sq.Select("LAST_INSERT_ID() AS id").
		RunWith(r.dbConn).
		QueryRowContext(ctx).
		Scan(&insertedID)
	if err != nil {
		logger.Errorf("error scan query select last id, error: %v", err)
		return nil, err
	}

	data.ID = insertedID
	data.CreatedAt = timeNow

	return &data, nil
}

func (r *customerRepository) Delete(ctx context.Context, id int64) error {
	panic("TODO")
}
