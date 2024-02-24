package repository

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/rezaig/dbo-service/internal/model"
)

type customerRepository struct {
	dbConn *sql.DB
}

func NewCustomerRepository(dbConn *sql.DB) model.CustomerRepository {
	return &customerRepository{dbConn: dbConn}
}

func (r *customerRepository) FindAll(ctx context.Context) ([]model.Customer, error) {
	rows, err := sq.Select("*").
		From("customer").
		RunWith(r.dbConn).
		QueryContext(ctx)
	if err != nil {
		log.Println("error select all, error: ", err)
		return nil, err
	}

	var results []model.Customer
	for rows.Next() {
		var result model.Customer
		err = rows.Scan(&result.ID, &result.Name)
		if err != nil {
			log.Println("error scanning, error: ", err)
			continue
		}
		results = append(results, result)
	}

	return results, nil
}
