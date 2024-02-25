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

type orderRepository struct {
	dbConn *sql.DB
}

func NewOrderRepository(dbConn *sql.DB) model.OrderRepository {
	return &orderRepository{dbConn: dbConn}
}

func (r *orderRepository) FindAllByCustomerID(ctx context.Context, customerID int64, params model.OrderParams) ([]model.Order, int64, error) {
	logger := log.WithFields(log.Fields{
		"func":       helper.GetFuncName(),
		"customerID": customerID,
		"params":     helper.Dump(params),
	})

	selectQ := sq.Select(
		"o.id AS order_id",
		"o.quantity",
		"o.shipping_address",
		"o.created_at",
		"c.id AS customer_id",
		"c.name AS customer_name",
		"p.id AS product_id",
		"p.name AS product_name").
		From("`order` o").
		Join("customer c ON c.id=o.customer_id").
		LeftJoin("product p ON p.id=o.product_id").
		Where(sq.Eq{"o.customer_id": customerID})

	if params.Keyword != "" {
		selectQ = selectQ.Where(sq.Like{"CAST(o.id AS CHAR)": fmt.Sprintf("%%%s%%", params.Keyword)})
	}

	log.Info(selectQ.ToSql())

	rows, err := selectQ.
		Limit(uint64(params.PerPage)).
		Offset(uint64(params.GetOffset())).
		RunWith(r.dbConn).
		QueryContext(ctx)
	if err != nil {
		logger.Errorf("error run query select all, error: %v", err)
		return nil, 0, err
	}

	var results []model.Order
	for rows.Next() {
		var result model.Order
		var (
			custID       int64
			customerName string
			productID    sql.NullInt64
			productName  sql.NullString
		)
		err = rows.Scan(
			&result.ID,
			&result.Quantity,
			&result.ShippingAddress,
			&result.CreatedAt,
			&custID,
			&customerName,
			&productID,
			&productName)
		if err != nil {
			logger.Errorf("error scanning query select all, error: %v", err)
			continue
		}
		result.Customer = model.Customer{
			ID:   custID,
			Name: customerName,
		}
		result.Product = model.Product{
			ID:   productID.Int64,
			Name: productName.String,
		}
		results = append(results, result)
	}

	selectCountQ := sq.Select("COUNT(o.id) AS total_items").
		From("`order` o").
		InnerJoin("customer c ON c.id=o.customer_id").
		LeftJoin("product p ON p.id=o.product_id").
		Where(sq.Eq{"customer_id": customerID})
	if params.Keyword != "" {
		selectQ = selectQ.Where(sq.Like{"CAST(o.id AS CHAR)": fmt.Sprintf("%%%s%%", params.Keyword)})
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

func (r *orderRepository) Insert(ctx context.Context, data model.Order) error {
	panic("TODO")
}
