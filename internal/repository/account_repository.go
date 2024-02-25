package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/rezaig/dbo-service/internal/helper"
	log "github.com/sirupsen/logrus"

	"github.com/rezaig/dbo-service/internal/model"
)

type accountRepository struct {
	dbConn *sql.DB
}

func NewAccountRepository(dbConn *sql.DB) model.AccountRepository {
	return &accountRepository{dbConn: dbConn}
}

func (r *accountRepository) Insert(ctx context.Context, data model.Account) error {
	timeNow := time.Now().UTC()
	_, err := sq.Insert("account").
		Columns("id", "username", "password", "created_at").
		Values(data.ID, data.Username, data.Password, timeNow).
		RunWith(r.dbConn).
		ExecContext(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"func": helper.GetFuncName(),
			"data": helper.Dump(data),
		}).Errorf("error exec query insert, error: %v", err)
	}

	return err
}
